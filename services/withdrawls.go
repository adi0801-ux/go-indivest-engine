package services

import (
	"encoding/json"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
)

func (p *MFService) CreateWithdrawl(createWithdrawls *models.CreateWithdrawls) (int, interface{}, error) {
	userDtls, err := p.ShowAccountRepo.ReadShowAccount(createWithdrawls.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.CreateWithdrawlAPI{}
	baseModel.Amount = createWithdrawls.Amount
	baseModel.FundCode = createWithdrawls.FundCode
	baseModel.AccountUuid = userDtls.AcntUuid
	baseModel.PartnerTransactionId = createWithdrawls.PartnerTransactionId
	response, err := p.TSAClient.SendPostRequest(constants.CreateWithdrawls, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.CreateWithdrawlAPIResponse
	//convert struct to []byte
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, nil
}
