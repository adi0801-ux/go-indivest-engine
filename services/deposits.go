package services

import (
	"encoding/json"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
	"net/url"
)

func (p *ServiceConfig) GetDeposits(getDeposits *models.GetDeposits) (int, interface{}, error) {
	baseModel := models.GetDepositsAPI{}
	baseModel.AccountUuid = getDeposits.AccountUuid
	params := url.Values{}
	params.Add("account_uuid", getDeposits.AccountUuid)
	response, err := p.TSAClient.SendGetRequest(constants.GetDeposits, params)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.GetDepositsAPIResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, err
}

func (p *ServiceConfig) ShowDeposits() (int, interface{}, error) {
	return 0, nil, nil
}

func (p *ServiceConfig) CreateDeposit(createDeposit *models.CreateDeposit) (int, interface{}, error) {
	baseModel := models.CreateDepositAPI{}
	baseModel.Amount = createDeposit.Amount
	baseModel.FundCode = createDeposit.FundCode
	baseModel.PaymentRedirectUrl = createDeposit.PaymentRedirectUrl
	baseModel.AccountUuid = createDeposit.AccountUuid
	baseModel.OnBoardingUuid = createDeposit.OnBoardingUuid
	baseModel.PartnerTransactionId = createDeposit.PartnerTransactionId

	response, err := p.TSAClient.SendPostRequest(constants.CreateDeposit, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.CreateDepositAPIResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, err
}
