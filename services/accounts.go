package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
	"net/url"
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
		UserId:   userInfo.UserId,
		Uuid:     userInfo.Uuid,
		AcntUuid: data.AcntUuid,
	}
	//create Db for show account
	err = p.ShowAccountRepo.CreateAccount(account)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, data, nil
}

func (p *MFService) ConnectWebhooks(webhooks *models.Webhook) (int, interface{}, error) {
	//verify payload
	// check if signature is hash of payload -->
	h := hmac.New(sha256.New, []byte(p.config.SecretKey))

	// Write Data to it
	b, err := json.Marshal(webhooks.Payload)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	h.Write([]byte(b))

	// Get result and encode as hexadecimal string
	sha := h.Sum(nil)

	if !hmac.Equal([]byte(webhooks.Signature), sha) {
		//	not from savvy ,  return
	}
	utils.Log.Info(webhooks.Event, webhooks.Payload)

	if webhooks.Event == constants.WebhooksCreateDeposits {
		_ = webhooks.Payload.(models.CreateDeposit)

	} else if webhooks.Event == constants.WebhooksCreateOnboardings {
		_ = webhooks.Payload.(models.CheckKYCUserAPI)

	} else if webhooks.Event == constants.WebhooksDepositsStatusUpdate {
		//which is depositStatusUpdate
		_ = webhooks.Payload.(models.Webhook)

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
