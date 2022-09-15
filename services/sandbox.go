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
		Active:     constants.DefaultSIPActiveSatus,
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

func (u *SandboxServiceConfig) GetActiveSIPWithTodaysDate() (*[]models.UserMFActiveSIP, error) {
	SipDate := utils.GetCurrentDate()

	SIPs, err := u.SandboxRep.FindAllUserSIPWithDate(SipDate)
	if err != nil {
		return nil, err
	}
	return SIPs, nil
}

func (u *SandboxServiceConfig) ProcessSIP() error {
	utils.Log.Infof("processing SIP's strarted")

	SIPs, err := u.GetActiveSIPWithTodaysDate()
	if err != nil {
		return err
	}

	for _, sip := range *SIPs {
		//	create a buying MF transaction
		userTransaction := &models.BuyMutualFund{
			InvestmentType: constants.SIP,
			SchemeCode:     sip.SchemeCode,
			Amount:         sip.SIPAmount,
			UserId:         sip.UserID,
		}

		wallet, err := u.SandboxRep.ReadUserWallet(userTransaction.UserId)
		if err != nil {
			utils.Log.Error("error fetching wallet ", err)
			continue
		}

		nav, err := u.GetNav(sip.SchemeCode)
		if err != nil {
			utils.Log.Error("error fetching nav ", err)
			continue
		}

		err = u.CreateBuyMfTransaction(userTransaction, wallet, nav, sip.SipID)
		if err != nil {
			utils.Log.Error("error processing transaction ", err)
			continue
		}

	}

	utils.Log.Info("processing SIP's completed")
	return nil

}
