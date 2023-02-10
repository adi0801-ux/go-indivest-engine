package services

import (
	"encoding/json"
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

	utils.Log.Info(webhooks.Event, webhooks.Payload)

	if webhooks.Event == constants.WebhooksCreateDeposits {
		webhookPayload := webhooks.Payload.(models.WebhookDepositsCreate)
		//onboardinguuid or uuid
		//onboardingObject, err := p.SavvyRepo.ReadOnboardingObjectByUUID(webhookPayload.Deposit.Uuid)
		depositObject, err := p.SavvyRepo.ReadDepositsByUUID(webhookPayload.Deposit.Uuid)
		if err != nil {
			utils.Log.Error(err)
			return http.StatusBadRequest, nil, err
		}
		depositObject.PaymentStatus = "Payment Initiated"
		err = p.SavvyRepo.CreateOrUpdateDeposit(depositObject)
		if err != nil {
			utils.Log.Error(err)
			return http.StatusBadRequest, nil, err
		}
		// will get this only once
		// get account by userId and amc id
		account := &models.ShowAccountDB{
			UserId:   depositObject.UserId,
			AmcId:    strconv.Itoa(webhookPayload.Deposit.Fund.AmcId),
			AcntUuid: webhookPayload.Deposit.AccountUuid,
		}

		_, err = p.SavvyRepo.ReadAccountWithAmcId(account.UserId, account.AmcId)
		if err != nil {
			if err.Error() == constants.UserNotFound {
				//	add the account
				err = p.SavvyRepo.CreateOrUpdateAccount(account)
				if err != nil {
					utils.Log.Error(err)
					return http.StatusBadRequest, nil, err
				}
			} else {
				utils.Log.Error(err)
				return http.StatusBadRequest, nil, err
			}
		}

		//	we need to update status for the transaction occured

	} else if webhooks.Event == constants.WebhooksCreateOnboardings {
		webhooks := webhooks.Payload.(models.WebhookOnboardingCreate)
		onboardingObject, err := p.SavvyRepo.ReadOnboardingObject(webhooks.Uuid)
		if err != nil {
			utils.Log.Error(err)
			return http.StatusBadRequest, nil, err
		}
		onboardingObject.OnboardingStatus = "Onboarding Object Created"
		err = p.SavvyRepo.UpdateOrCreateOnboardingObject(onboardingObject)
		if err != nil {
			utils.Log.Error(err)
			return http.StatusBadRequest, nil, err
		}

	} else if webhooks.Event == constants.WebhooksDepositsStatusUpdate {
		//which is depositStatusUpdate
		webhookPayload := webhooks.Payload.(models.WebhookDepositsCreate)

		depositObject, err := p.SavvyRepo.ReadDepositsByUUID(webhookPayload.Deposit.Uuid)
		if err != nil {
			utils.Log.Error(err)
			return http.StatusBadRequest, nil, err
		}
		if webhookPayload.Deposit.Status == "created" {
			depositObject.PaymentStatus = "Payment Transaction Success"
		} else {
			webhookPayload.Deposit.Status = "Payment Transaction Filed"
		}

		err = p.SavvyRepo.CreateOrUpdateDeposit(depositObject)
		if err != nil {
			utils.Log.Error(err)
			return http.StatusBadRequest, nil, err
		}

	} else if webhooks.Event == constants.WebhooksCreateAccounts {
		//which is the createAccount api? is it the AddBankAccount under Onboarding
		//depositsPayload := webhooks.Payload.(models.createa)
		utils.Log.Info(webhooks.Event, webhooks.Payload)
	} else if webhooks.Event == constants.WebhooksCreateWithdrawals {
		//on  what basis do i select model of withdrawal from createwithdrawal, createwithdrawalapi, response
		_ = webhooks.Payload.(models.CreateWithdrawals)

	} else if webhooks.Event == constants.WebhooksWithdrawalStatusUpdate {
		_ = webhooks.Payload.(models.Webhook)
	}
	return http.StatusOK, nil, nil
}
