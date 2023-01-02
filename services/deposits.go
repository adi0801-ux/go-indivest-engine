package services

import (
	"encoding/json"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
)

func(p *ServiceConfig)GetDeposits(getDeposits *models.GetDeposits)(int, interface{}, error{
	baseModel := models.GetDepositsAPI{}
	baseModel.AccountUuid = getDeposits.AccountUuid

	response, err := p.TSAClient.SendPostRequest(constants.GetDeposits,&baseModel)
	if err!=nil{
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



