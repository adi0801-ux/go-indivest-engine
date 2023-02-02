package services

import (
	"encoding/json"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
	"net/url"
)

//Accounts API
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
	err = p.ShowAccountRepo.CreateShowAccount(account)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, data, nil
}
