package services

import "indivest-engine/models"

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
