package services

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
)

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
