package services

import (
	"encoding/json"
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
	"net/url"
	"strconv"
)

// Accounts API
func (p *MFService) ShowAccounts(userIdDtls *models.ShowAccount) (int, interface{}, error) {
	userInfo, err := p.SavvyRepo.ReadOnboardingObject(userIdDtls.UserId)
	//baseModel := models.ShowAccountAPI{}
	//data, err := p.SavvyRepo.ReadAllAccounts()
	params := url.Values{}

	response, err := p.TSAClient.SendGetRequest(constants.GenerateShowAccountsURL(userInfo.Uuid), params)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	//response to get AcntUuid
	var data models.ShowAccountAPIResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	account := &models.ShowAccountDB{
		UserId: userInfo.UserId,
		//Uuid:     userInfo.Uuid,
		AcntUuid: data.AcntUuid,
	}
	//create Db for show account
	err = p.SavvyRepo.CreateAccount(account)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, data, nil
}

func (p *MFService) ConnectWebhooks(webhooks *models.Webhook) (int, interface{}, error) {

	//fmt.Print(webhooks)
	//fmt.Print("webhook connected")
	if webhooks.Event == constants.WebhooksCreateDeposits {
		//fmt.Print("webhook hit")
		err := p.depositCreateWebhook(webhooks.Payload)
		if err != nil {
			utils.Log.Info(err)
			return http.StatusBadRequest, nil, err
		}
	} else if webhooks.Event == constants.WebhooksStatusUpdateDeposits {
		err := p.depositStatusUpdateWebhook(webhooks.Payload)
		if err != nil {
			return http.StatusBadRequest, nil, err
		}
	} else if webhooks.Event == constants.WebhooksCreateOnboardings {
		err := p.onboardingCreateWebhook(webhooks.Payload)
		if err != nil {
			return http.StatusBadRequest, nil, err
		}

		//} else if webhooks.Event == constants.WebhooksUpdateOnboarding {
		//	err := p.onboardingUpdateWebhook(webhooks.Payload)
		//	if err != nil {
		//		return http.StatusBadRequest, nil, err
		//	}
	} else if webhooks.Event == constants.WebhooksCreateAccounts {
		err := p.accountWebhook(webhooks.Payload)
		if err != nil {
			return http.StatusBadRequest, nil, err
		}
	} else if webhooks.Event == constants.WebhooksCreateWithdrawals {
		err := p.withdrawCreateWebhook(webhooks.Payload)
		if err != nil {
			return http.StatusBadRequest, nil, err
		}
	} else if webhooks.Event == constants.WebhooksStatusUpdateWithdrawal {
		err := p.withdrawStatusUpdateWebhook(webhooks.Payload)
		if err != nil {
			return http.StatusBadRequest, nil, err
		}
	} else if webhooks.Event == constants.WebhooksSipCreated {
		err := p.sipCreateWebhook(webhooks.Payload)
		if err != nil {
			return http.StatusBadRequest, nil, err
		}
	}
	return http.StatusOK, nil, nil
}
func (p *MFService) sipCreateWebhook(webhookPayload interface{}) error {
	var sipCreatePayload models.WebhookSipCreate
	err := utils.Transcode(webhookPayload, &sipCreatePayload)
	if err != nil {
		//cannot convert to struct
		utils.Log.Error(err)
		return err
	}
	sipObject, err := p.SavvyRepo.ReadSipUuid(sipCreatePayload.Sip.Uuid)
	if err != nil {
		utils.Log.Info(err)
		return err
	}
	sipObject.SipStatus = constants.SipCreated
	fmt.Println(sipObject.SipStatus)
	err = p.SavvyRepo.UpdateSip(sipObject)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	return nil
}
func (p *MFService) onboardingCreateWebhook(webhookPayload interface{}) error {
	var onboardingPayload models.WebhookOnboardingCreate

	err := utils.Transcode(webhookPayload, &onboardingPayload)
	if err != nil {
		//cannot convert to struct
		utils.Log.Error(err)
		return err
	}
	//fmt.Println(onboardingPayload)
	onboardingObject, err := p.SavvyRepo.ReadOnboardingObjectByUUID(onboardingPayload.Onboarding.Uuid)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	fmt.Println(onboardingObject)
	onboardingObject.OnboardingStatus = "Onboarding Object Created"
	err = p.SavvyRepo.UpdateOrCreateOnboardingObjectUuid(onboardingObject)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	return nil
}

//	func (p *MFService) onboardingUpdateWebhook(webhookPayload interface{}) error {
//		var onboardingPayload models.WebhookOnboardingUpdate
//
//		err := utils.Transcode(webhookPayload, &onboardingPayload)
//		if err != nil {
//			//cannot convert to struct
//			utils.Log.Error(err)
//			return err
//		}
//		//fmt.Println(onboardingPayload)
//		onboardingObject, err := p.SavvyRepo.ReadOnboardingObjectByUUID(onboardingPayload.Onboarding.Uuid)
//		if err != nil {
//			utils.Log.Error(err)
//			return err
//		}
//		fmt.Println(onboardingObject)
//		onboardingObject.OnboardingStatus = "Onboarding Object Created"
//		err = p.SavvyRepo.UpdateOrCreateOnboardingObjectUuid(onboardingObject)
//		if err != nil {
//			utils.Log.Error(err)
//			return err
//		}
//		return nil
//	}
func (p *MFService) withdrawStatusUpdateWebhook(webhookPayload interface{}) error {
	var onboardingPayload models.WebhookWithdrawCreate

	err := utils.Transcode(webhookPayload, &onboardingPayload)
	if err != nil {
		//cannot convert to struct
		utils.Log.Error(err)
		return err
	}

	withdrawalObject, err := p.SavvyRepo.ReadWithdrawalUuid(onboardingPayload.Withdrawal.Uuid)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	withdrawalObject.WithdrawalStatus = "Withdrawal Completed"
	err = p.SavvyRepo.UpdateWithdrawalUuid(withdrawalObject)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	return nil
}

func (p *MFService) withdrawCreateWebhook(webhookPayload interface{}) error {
	var accountPayload models.WebhookWithdrawCreate

	err := utils.Transcode(webhookPayload, &accountPayload)
	if err != nil {
		//cannot convert to struct
		utils.Log.Error(err)
		return err
	}
	withdrawalObject, err := p.SavvyRepo.ReadWithdrawalUuid(accountPayload.Withdrawal.Uuid)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	withdrawalObject.WithdrawalStatus = "Withdrawal Initiated"
	withdrawalObject.AmcId = strconv.Itoa(accountPayload.Withdrawal.Fund.AmcId)
	err = p.SavvyRepo.UpdateWithdrawalUuid(withdrawalObject)
	if err != nil {
		utils.Log.Info(err)
		return err
	}
	return nil
}

func (p *MFService) depositStatusUpdateWebhook(webhookPayload interface{}) error {
	var depositsPayload models.WebhookDepositsCreate
	err := utils.Transcode(webhookPayload, &depositsPayload)
	if err != nil {
		//cannot convert to struct
		utils.Log.Error(err)
		return err
	}
	depositObject, err := p.SavvyRepo.ReadDepositsByUUID(depositsPayload.Deposit.Uuid)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	if depositsPayload.Deposit.Status == "payment_made" {
		depositObject.PaymentStatus = "Payment Transaction Success"
	} else {
		depositObject.PaymentStatus = "Payment Transaction Failed"
	}
	err = p.SavvyRepo.CreateOrUpdateDepositUuid(depositObject)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	return nil
}

func (p *MFService) depositCreateWebhook(webhookPayload interface{}) error {
	var depositCreate models.WebhookDepositsCreate
	//fmt.Print("create Deposit webhook")
	err := utils.Transcode(webhookPayload, &depositCreate)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	sip := &models.CreateSipDb{
		MonthlySipStatus: "Monthly SIP amount deducted",
	}
	if depositCreate.Deposit.SipUuid != "" {
		err = p.SavvyRepo.UpdateSip(sip)
		if err != nil {
			utils.Log.Error(err)
			return err
		}
	}
	//onboardinguuid or uuid
	//onboardingObject, err := p.SavvyRepo.ReadOnboardingObjectByUUID(webhookPayload.Deposit.Uuid)
	depositObject, err := p.SavvyRepo.ReadDepositsByUUID(depositCreate.Deposit.Uuid)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	depositObject.PaymentStatus = "Payment Initiated"
	fmt.Println(depositObject.PaymentStatus)
	err = p.SavvyRepo.CreateOrUpdateDepositUuid(depositObject)
	if err != nil {
		fmt.Print(err)
		utils.Log.Error(err)
		return err
	}
	// will get this only once
	// get account by userId and amc id
	account := &models.ShowAccountDB{
		UserId:   depositObject.UserId,
		AmcId:    strconv.Itoa(depositCreate.Deposit.Fund.AmcId),
		AcntUuid: depositCreate.Deposit.AccountUuid,
	}

	_, err = p.SavvyRepo.ReadAccountWithAmcId(account.UserId, account.AmcId)
	if err != nil {
		if err.Error() == constants.UserNotFound {
			//	add the account
			err = p.SavvyRepo.CreateOrUpdateAccount(account)
			if err != nil {
				utils.Log.Error(err)
				return err
			}
		} else {
			utils.Log.Error(err)
			return err
		}
	}
	return nil
}

func (p *MFService) accountWebhook(webhookPayload interface{}) error {
	var accountPayload models.WebhookAccountCreate

	err := utils.Transcode(webhookPayload, &accountPayload)
	if err != nil {
		//cannot convert to struct
		utils.Log.Error(err)
		return err
	}

	onboardingObject, err := p.SavvyRepo.ReadOnboardingObjectByUUID(accountPayload.Account.OnboardingUuid)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	//fmt.Print(accountPayload.Account.AmcCode)
	amcInfo, err := p.SavvyRepo.ReadFundHouseDetailsWithAmcCode(accountPayload.Account.AmcCode)
	fmt.Print(amcInfo)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	//fmt.Print(amcInfo.AMCID)
	//create model
	account := &models.ShowAccountDB{UserId: onboardingObject.UserId, AmcId: strconv.Itoa(amcInfo.AMCID), AcntUuid: accountPayload.Account.Uuid}
	err = p.SavvyRepo.CreateOrUpdateAccount(account)
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	return nil
}
