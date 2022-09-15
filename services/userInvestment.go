package services

import (
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"time"
)

func (u *SandboxServiceConfig) InvestmentAnalysis(userID string) []models.InvestmentAnalysis {
	//	get all transactions from user

	allTransactions, err := u.GetUserAllTransactions(userID)
	if err != nil {
		return []models.InvestmentAnalysis{}
	}

	//	group by scheme code
	//	group by buy/sell
	transactions := make(map[string]map[string][]models.UserMFTransactions)
	for _, transaction := range allTransactions {
		if transactions[transaction.SchemeCode] == nil {
			newMap := make(map[string][]models.UserMFTransactions)

			newMap["BUY"] = []models.UserMFTransactions{}
			newMap["SELL"] = []models.UserMFTransactions{}
			transactions[transaction.SchemeCode] = newMap
		}
		transactions[transaction.SchemeCode][transaction.TransactionType] = append(transactions[transaction.SchemeCode][transaction.TransactionType], transaction)

	}

	var investmentAnalysis []models.InvestmentAnalysis
	for s, m := range transactions {

		investmentAnalysis = append(investmentAnalysis, u.GetTransactionPerSide(s, m))
	}
	return investmentAnalysis
}

func (u *SandboxServiceConfig) GetTransactionPerSide(schemeCode string, transaction map[string][]models.UserMFTransactions) (t models.InvestmentAnalysis) {
	//	 for buy
	buyTransactions := transaction["BUY"]
	buyAmount, buyUnits := GetTotalAmounts(buyTransactions)

	//for sell
	sellTransactions := transaction["SELL"]
	sellAmount, sellUnits := GetTotalAmounts(sellTransactions)

	currentNav, _ := u.GetNav(schemeCode)

	t.Units = utils.RoundOfTo2Decimal(buyUnits - sellUnits)
	t.InvestedAmount = utils.RoundOfTo2Decimal(buyAmount - sellAmount)
	t.CurrentWorth = utils.RoundOfTo2Decimal(t.Units * currentNav)
	t.PNL = utils.RoundOfTo2Decimal(t.CurrentWorth - t.InvestedAmount)
	t.PNLPercentage = utils.RoundOfTo2Decimal((t.PNL / t.InvestedAmount) * 100)
	t.SchemeCode = schemeCode

	return
}

func GetTotalAmounts(transaction []models.UserMFTransactions) (float64, float64) {

	totalUnits := 0.0
	totalAmount := 0.0
	for _, mfTransaction := range transaction {

		totalUnits = totalUnits + mfTransaction.FundUnits
		totalAmount = totalAmount + mfTransaction.INRAmount
	}

	return totalAmount, totalUnits

}

func (u *SandboxServiceConfig) InvestmentPanel(userID string) (models.UserMfInvestmentPanel, error) {
	todaysReport := u.InvestmentAnalysis(userID)
	investmentResult := models.UserMfInvestmentPanel{}
	for _, transaction := range todaysReport {
		investmentResult.TotalInvestment = investmentResult.TotalInvestment + transaction.InvestedAmount

		investmentResult.CurrentWorth = investmentResult.CurrentWorth + transaction.CurrentWorth
	}

	return investmentResult, nil
}

func (u *SandboxServiceConfig) UserMfActivity(userID string) (map[string][]models.UserMfActivity, error) {
	previousReport, err := u.SandboxRep.ReadAllMFDailyReport(userID, constants.DefaultActivityDayLimit)
	if err != nil {
		return map[string][]models.UserMfActivity{}, err
	}
	todaysReport := u.InvestmentAnalysis(userID)

	//	 make map
	mfActivity := make(map[string][]models.UserMfActivity)

	for _, transaction := range *previousReport {
		if mfActivity[transaction.SchemeCode] == nil {
			activity := models.UserMfActivity{
				CurrentWorth: transaction.InvestmentWorth,
				Date:         transaction.CreatedAt.String(),
			}
			mfActivity[transaction.SchemeCode] = []models.UserMfActivity{activity}
		} else {
			activity := models.UserMfActivity{
				CurrentWorth: transaction.InvestmentWorth,
				Date:         transaction.CreatedAt.String(),
			}

			mfActivity[transaction.SchemeCode] = append(mfActivity[transaction.SchemeCode], activity)
		}
	}
	for _, transaction := range todaysReport {
		if mfActivity[transaction.SchemeCode] == nil {
			activity := models.UserMfActivity{
				CurrentWorth: transaction.CurrentWorth,
				Date:         time.Now().Local().String(),
			}

			mfActivity[transaction.SchemeCode] = []models.UserMfActivity{activity}
		} else {
			activity := models.UserMfActivity{
				CurrentWorth: transaction.CurrentWorth,
				Date:         time.Now().Local().String(),
			}
			mfActivity[transaction.SchemeCode] = append(mfActivity[transaction.SchemeCode], activity)
		}
	}
	return mfActivity, nil
}

func (u *SandboxServiceConfig) DailyReport() error {
	//	for every user in sandbox --> from wallet
	userWallets, err := u.SandboxRep.ReadAllUserWallets()
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	for _, wallet := range *userWallets {

		investmentAnalysis := u.InvestmentAnalysis(wallet.UserID)

		for _, investments := range investmentAnalysis {
			//	store current worth with scheme code in db
			report := models.UserMFDailyReport{
				UserID:          wallet.UserID,
				SchemeCode:      investments.SchemeCode,
				InvestmentWorth: investments.CurrentWorth,
			}

			err := u.SandboxRep.CreateMFDailyReport(&report)
			if err != nil {
				utils.Log.Error(err)
			}
		}
	}
	return nil
}
