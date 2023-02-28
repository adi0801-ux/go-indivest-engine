package services

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (p *RiskCalculatorService) AddLanguage(language *models.UserBasicDetailsLanguage) error {

	var userDetails models.UserDetails
	userDetails.UserID = language.UserId
	userDetails.Language = language.Language

	err := p.UserRepo.CreateOrUpdateUserDetails(&userDetails)

	return err

}

func (p *RiskCalculatorService) AddIncome(Income *models.UserBasicDetailsIncome) error {

	var userDetails models.UserDetails

	userDetails.UserID = Income.UserId
	userDetails.Age = Income.Age
	userDetails.GrossMonthlyIncome = utils.RoundOfTo2Decimal(Income.Income)
	fmt.Println(utils.RoundOfTo2Decimal(Income.Income))
	fmt.Println(utils.RoundOfTo2Decimal(12.3456))
	userDetails.Profession = Income.Profession
	userDetails.UserExpertise = Income.UserExpertise
	err := p.UserRepo.CreateOrUpdateUserDetails(&userDetails)
	if err != nil {
		utils.Log.Error(err)
		return err
	}

	return err
}

func (p *RiskCalculatorService) AddExpenses(Expenses *models.UserBasicDetailsExpenses) (calcResp models.CalculationResponse, err error) {

	userDetails, err := p.UserRepo.ReadUserDetails(Expenses.UserId)
	if err != nil {
		return calcResp, err
	}

	userDetails.UserID = Expenses.UserId
	userDetails.MonthlyEssentialExpense = utils.RoundOfTo2Decimal(Expenses.MonthlyEssentialExpense)
	userDetails.MonthlyNonEssentialExpense = utils.RoundOfTo2Decimal(Expenses.MonthlyNonEssentialExpense)
	userDetails.MonthlySavings = userDetails.GrossMonthlyIncome - userDetails.MonthlyNonEssentialExpense - userDetails.MonthlyEssentialExpense
	userDetails.MonthlyInvestments = utils.RoundOfTo2Decimal(Expenses.MonthlyInvestments)
	userDetails.MonthlyInvestibleSurplus = userDetails.MonthlySavings - userDetails.MonthlyInvestments

	if userDetails.MonthlySavings < 0 {
		return calcResp, fmt.Errorf("monthly_income is more or less than expenses (not possible)")
	}

	if userDetails.MonthlyInvestibleSurplus < 0 {
		return calcResp, fmt.Errorf("monthly_savings is more than monthly_investments (not possible)")
	}

	err = p.UserRepo.UpdateUserDetails(userDetails)
	if err != nil {
		return calcResp, err
	}

	user, err := p.UserRepo.ReadUser(Expenses.UserId)
	if err != nil {
		utils.Log.Error(err)
		return calcResp, err
	}
	user.ProfileStatus = "2"

	err = p.UserRepo.UpdateOrCreateUser(user)
	if err != nil {
		utils.Log.Error(err)
		return calcResp, err
	}
	return p.CalculateUserInformation(userDetails)
}

func (p *RiskCalculatorService) GetUserInformation(UserId string) (calcResp models.CalculationResponse, err error) {
	userDetails, err := p.UserRepo.ReadUserDetails(UserId)
	if err != nil {
		return calcResp, err
	}

	return p.CalculateUserInformation(userDetails)

}

func (p *RiskCalculatorService) CalculateUserInformation(userDetails *models.UserDetails) (calcResp models.CalculationResponse, err error) {

	surplus, err := p.CalculateRecommendedInvestibleSurplus(userDetails)
	if err != nil {
		return calcResp, err
	}
	calcResp.InvestibleSurplus = surplus

	stats, err := p.CalculateCurrentPercentStats(userDetails)
	if err != nil {
		return calcResp, err
	}
	calcResp.CurrentPercentStats = stats

	signal := p.CalculateHealthSignal(&stats)
	calcResp.HealthSignal = signal

	idealStats, err := p.CalculateIdealFinancialProfile(userDetails)
	if err != nil {
		return calcResp, err
	}
	calcResp.IdealPercentStats = idealStats

	return

}

func (p *RiskCalculatorService) CalculateRecommendedInvestibleSurplus(userDetails *models.UserDetails) (models.InvestibleSurplus, error) {

	var rcmInvestibleFund models.InvestibleSurplus

	if userDetails.MonthlyInvestibleSurplus <= 0 {
		rcmInvestibleFund.EmergencyFund = utils.RoundOfTo2Decimal(0)
	} else if userDetails.MonthlyInvestibleSurplus*constants.RecommendedInvestibleFund < 1000 {
		rcmInvestibleFund.EmergencyFund = utils.RoundOfTo2Decimal(userDetails.MonthlyInvestibleSurplus)
	} else if userDetails.MonthlyInvestibleSurplus*constants.RecommendedInvestibleFund > 1000 {
		rcmInvestibleFund.EmergencyFund = utils.RoundOfTo2Decimal(constants.RecommendedEmergencyFund * userDetails.MonthlyInvestibleSurplus)
	}

	if userDetails.MonthlyInvestibleSurplus*constants.RecommendedInvestibleFund >= 1000 {
		rcmInvestibleFund.InvestibleFund = utils.RoundOfTo2Decimal(userDetails.MonthlyInvestibleSurplus * constants.RecommendedInvestibleFund)
	} else {
		rcmInvestibleFund.InvestibleFund = utils.RoundOfTo2Decimal(0)
	}

	return rcmInvestibleFund, nil
}

func (p *RiskCalculatorService) CalculateCurrentPercentStats(userDetails *models.UserDetails) (models.CurrentPercentStats, error) {

	var currentStats models.CurrentPercentStats

	currentStats.EssentialExpenses = utils.RoundOfTo2Decimal(userDetails.MonthlyEssentialExpense / userDetails.GrossMonthlyIncome * 100)

	currentStats.NonEssentialExpenses = utils.RoundOfTo2Decimal(userDetails.MonthlyNonEssentialExpense / userDetails.GrossMonthlyIncome * 100)

	currentStats.Savings = utils.RoundOfTo2Decimal(userDetails.MonthlySavings / userDetails.GrossMonthlyIncome * 100)

	return currentStats, nil
}

func (p *RiskCalculatorService) CalculateHealthSignal(currentStats *models.CurrentPercentStats) string {

	var count int = 0

	if currentStats.EssentialExpenses <= 50 {
		count += 1
	}
	if currentStats.NonEssentialExpenses <= 30 {
		count += 1
	}
	if currentStats.Savings > 20 {
		count += 1
	}

	switch count {
	case 0, 1:
		return constants.HealthSignalRed
	case 2:
		return constants.HealthSignalAmber
	case 3:
		return constants.HealthSignalGreen
	}

	return constants.HealthSignalAmber
}

func (p *RiskCalculatorService) CalculateIdealFinancialProfile(userDetails *models.UserDetails) (models.IdealPercentStats, error) {

	var idealStats models.IdealPercentStats

	idealStats.EssentialExpenses = utils.RoundOfTo2Decimal(userDetails.GrossMonthlyIncome * constants.IdealMonthlyEssentialExpense / 100)

	idealStats.NonEssentialExpenses = utils.RoundOfTo2Decimal(userDetails.GrossMonthlyIncome * constants.IdealMonthlyNonEssentialExpense / 100)

	idealStats.Savings = utils.RoundOfTo2Decimal(userDetails.GrossMonthlyIncome * constants.IdealMonthlySavings / 100)

	return idealStats, nil

}
