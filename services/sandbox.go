package services

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"strconv"
	"time"
)

func (u *SandboxServiceConfig) BuyMutualFund(userTransaction *models.BuyMutualFund) error {

	//validate scheme code
	//add redis(assume validated )

	// fetch wallet
	var wallet *models.UserWallet
	var err error

	nav, err := u.GetNav(userTransaction.SchemeCode)
	if err != nil {
		return err
	}
	if nav == 0 {
		return fmt.Errorf(constants.SchemeCodeIsInvalid)
	}

	wallet, err = u.SandboxRep.ReadUserWallet(userTransaction.UserId)
	if err != nil {
		if err.Error() == constants.UserNotFound {
			wallet, err = u.SandboxRep.CreateGetUserWallet(userTransaction.UserId)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	if wallet.INR < userTransaction.Amount {
		return fmt.Errorf(constants.WalletAmountIsLow)
	}

	if userTransaction.InvestmentType == constants.SIP {
		return u.StartSIP(userTransaction, wallet, nav)
	} else {
		return u.OneTimePayment(userTransaction, wallet, nav)
	}

}

func (u *SandboxServiceConfig) StartSIP(userTransaction *models.BuyMutualFund, w *models.UserWallet, nav float64) error {

	SipDate := utils.GetCurrentDate()
	SidID := utils.GenerateSipID()

	newSip := &models.UserMFActiveSIP{
		UserID:     userTransaction.UserId,
		SipID:      SidID,
		SchemeCode: userTransaction.SchemeCode,
		SIPAmount:  userTransaction.Amount,
		SIPDate:    SipDate,
	}

	err := u.SandboxRep.CreateUserSIP(newSip)
	if err != nil {
		//error creating SIP
		return err
	}

	//create transaction
	err = u.CreateBuyMfTransaction(userTransaction, w, nav, SidID)

	return err
}

func (u *SandboxServiceConfig) OneTimePayment(userTransaction *models.BuyMutualFund, w *models.UserWallet, nav float64) error {
	err := u.CreateBuyMfTransaction(userTransaction, w, nav, "")
	return err
}

func (u *SandboxServiceConfig) CreateBuyMfTransaction(userTransaction *models.BuyMutualFund, w *models.UserWallet, nav float64, SipId string) error {
	unitsAllocated := userTransaction.Amount / nav

	//deduct wallet balance
	w.INR = w.INR - userTransaction.Amount

	err := u.SandboxRep.UpdateUserWallet(w)
	if err != nil {
		return err
	}

	transaction := &models.UserMFTransactions{
		TransactionID:   utils.GenerateTransactionID(),
		SipID:           SipId,
		UserID:          userTransaction.UserId,
		SchemeCode:      userTransaction.SchemeCode,
		FundUnits:       unitsAllocated,
		NAV:             nav,
		INRAmount:       userTransaction.Amount,
		TransactionType: "BUY",
		InvestmentType:  userTransaction.InvestmentType,
	}

	err = u.SandboxRep.CreateMFTransaction(transaction)
	if err != nil {
		return err
	}

	// add to holdings
	err = u.AddToHoldings(userTransaction, unitsAllocated)

	return err
}

func (u *SandboxServiceConfig) AddToHoldings(userTransaction *models.BuyMutualFund, unitsAllocated float64) error {
	//get holdings

	holdings, err := u.SandboxRep.ReadUserHolding(userTransaction.UserId, userTransaction.SchemeCode)
	if err != nil {
		holdings = &models.UserMFHoldings{
			UserID:     userTransaction.UserId,
			SchemeCode: userTransaction.SchemeCode,
			FundUnits:  unitsAllocated,
			UpdatedAt:  time.Time{},
		}
		err := u.SandboxRep.UpdateOrCreateUserHoldings(holdings)
		if err != nil {
			return err
		}

	} else {
		holdings.FundUnits = holdings.FundUnits + unitsAllocated
		err := u.SandboxRep.UpdateOrCreateUserHoldings(holdings)
		if err != nil {
			return err
		}

	}
	return nil

}

func (u *SandboxServiceConfig) RedeemMutualFund(userTransaction *models.RedeemMutualFund) error {
	//	 amount and scheme code to be given
	//	fetch current nav for scheme code
	nav, err := u.GetNav(userTransaction.SchemeCode)
	if err != nil {
		return err
	}

	wallet, err := u.SandboxRep.ReadUserWallet(userTransaction.UserId)
	if err != nil {
		return err
	}
	// create transaction
	err = u.CreateRedeemMfTransaction(userTransaction, wallet, nav)
	if err != nil {
		return err
	}

	return nil
}

func (u *SandboxServiceConfig) CreateRedeemMfTransaction(userTransaction *models.RedeemMutualFund, w *models.UserWallet, nav float64) error {
	//	deduct holding
	unitsDeducted := userTransaction.Amount / nav

	err := u.DeductHoldings(userTransaction, unitsDeducted)
	if err != nil {
		return err
	}

	//	create a transaction record
	transaction := &models.UserMFTransactions{
		TransactionID:   utils.GenerateTransactionID(),
		SipID:           "",
		UserID:          userTransaction.UserId,
		SchemeCode:      userTransaction.SchemeCode,
		FundUnits:       unitsDeducted,
		NAV:             nav,
		INRAmount:       userTransaction.Amount,
		TransactionType: "SELL",
		InvestmentType:  "",
	}

	err = u.SandboxRep.CreateMFTransaction(transaction)
	if err != nil {
		return err
	}

	// add to wallet and updated
	//	get number of units to be deducted from holding if exists
	//	add the amount to be added wallet

	w.INR = w.INR + userTransaction.Amount
	err = u.SandboxRep.UpdateUserWallet(w)
	if err != nil {
		return err
	}

	return nil

}

func (u *SandboxServiceConfig) GetNav(schemeCode string) (float64, error) {
	value, err := u.RedisRep.GetKeyValue("nav_" + schemeCode)
	if err != nil {
		return 0, err
	}
	nav, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	return nav, nil
}

func (u *SandboxServiceConfig) DeductHoldings(userTransaction *models.RedeemMutualFund, unitsDeducted float64) error {

	holdings, err := u.SandboxRep.ReadUserHolding(userTransaction.UserId, userTransaction.SchemeCode)
	if err != nil {
		return err
	}
	if holdings.FundUnits < unitsDeducted {
		return fmt.Errorf(constants.UnitsAmountIsLow)
	}

	holdings.FundUnits = holdings.FundUnits - unitsDeducted

	//save in db
	err = u.SandboxRep.UpdateOrCreateUserHoldings(holdings)
	if err != nil {
		return err
	}
	return nil
}

func (u *SandboxServiceConfig) GetUserHolding(userId string, schemeCode string) (holdings models.Holdings, err error) {
	userHolding, err := u.SandboxRep.ReadUserHolding(userId, schemeCode)
	if err != nil {
		return holdings, err
	}

	// get nav -- > change response structure
	nav, err := u.GetNav(schemeCode)
	if err != nil {
		return holdings, err
	}

	holdings.Units = userHolding.FundUnits
	holdings.SchemeCode = userHolding.SchemeCode
	holdings.CurrentValue = userHolding.FundUnits * nav

	return holdings, nil
}

func (u *SandboxServiceConfig) GetAllUserHoldings(userId string) (holdings []models.Holdings, err error) {
	userHoldings, err := u.SandboxRep.ReadUserHoldings(userId)
	if err != nil {
		return holdings, err
	}

	// get nav -- > change response structure
	for _, mfHoldings := range *userHoldings {

		//get nav
		nav, err := u.GetNav(mfHoldings.SchemeCode)
		if err != nil {
			return holdings, err
		}

		holding := &models.Holdings{
			SchemeCode:   mfHoldings.SchemeCode,
			Units:        mfHoldings.FundUnits,
			CurrentValue: mfHoldings.FundUnits * nav,
		}
		holdings = append(holdings, *holding)
	}

	return holdings, nil
}

func (u *SandboxServiceConfig) GetUserWallet(userId string) (models.UserWallet, error) {
	wallet, err := u.SandboxRep.ReadUserWallet(userId)
	if err != nil {
		return models.UserWallet{}, err
	}

	return *wallet, nil
}

func (u *SandboxServiceConfig) GetUserAllTransactions(userId string) ([]models.UserMFTransactions, error) {
	transactions, err := u.SandboxRep.ReadAllMFTransactions(userId)
	if err != nil {
		return []models.UserMFTransactions{}, err
	}
	return *transactions, nil
}

func (u *SandboxServiceConfig) GetUserTransactions(userId string, schemeCode string) ([]models.UserMFTransactions, error) {
	transactions, err := u.SandboxRep.ReadMFTransaction(userId, schemeCode)
	if err != nil {
		return []models.UserMFTransactions{}, err
	}
	return *transactions, nil
}
