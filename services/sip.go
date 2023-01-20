package services

import (
	"encoding/json"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
	"net/url"
)

func (p *MFService) GetSip(getSip *models.GetSip) (int, interface{}, error) {
	baseModel := models.GetSipAPI{}
	baseModel.AccountUuid = getSip.AccountUuid
	params := url.Values{}
	params.Add("account_uuid", getSip.AccountUuid)
	response, err := p.TSAClient.SendGetRequest(constants.GetSip, params)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.GetSipAPIResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, err
}

func (p *MFService) ShowSip() (int, interface{}, error) {
	return http.StatusOK, nil, nil
}

func (p *MFService) CreateSip(createSip *models.CreateSip) (int, interface{}, error) {
	baseModel := models.CreateSipAPI{}
	baseModel.Amount = createSip.Amount
	baseModel.FundCode = createSip.FundCode
	baseModel.AccountUuid = createSip.AccountUuid
	baseModel.OnBoardingUuid = createSip.OnBoardingUuid
	baseModel.PartnerTransactionId = createSip.PartnerTransactionId
	baseModel.StartDate = createSip.StartDate
	baseModel.EndDate = createSip.EndDate
	baseModel.Frequency = createSip.Frequency
	baseModel.MandateRedirectUrl = createSip.MandateRedirectUrl

	response, err := p.TSAClient.SendPostRequest(constants.CreateSip, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.CreateSipAPIResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, err
}
