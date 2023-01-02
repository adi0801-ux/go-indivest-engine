package services

import (
	"encoding/json"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
)

func (p *ServiceConfig) GetSip(getSip *models.GetSip) (int, interface{}, error) {
	baseModel := models.GetSipAPI{}
	baseModel.AccountUuid = getSip.AccountUuid
	response, err := p.TSAClient.SendPostRequest(constants.GetSip, &baseModel)
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

func (p *ServiceConfig) ShowSip() (int, interface{}, error) {
	return http.StatusOK, nil, nil
}
