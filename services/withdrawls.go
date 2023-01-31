package services

import (
	"encoding/json"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
)

func (p *MFService) VerifyWithdrawalOtp(verifyOtp *models.VerifyWithdrawalOtp) (int, interface{}, error) {
	userDtls, err := p.ShowAccountRepo.ReadShowAccount(verifyOtp.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.VerifyWithdrawalOtpAPI{}
	baseModel.Otp = verifyOtp.Otp
	response, err := p.TSAClient.SendPostRequest(constants.GenerateVerifyWithdrawalOtpUrl(userDtls.AcntUuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.VerifyWithdrawOtpAPIResponse
	//convert struct to []byte
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, nil

}
func (p *MFService) CreateWithdrawal(createWithdrawals *models.CreateWithdrawals) (int, interface{}, error) {
	userDtls, err := p.ShowAccountRepo.ReadShowAccount(createWithdrawals.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.CreateWithdrawalAPI{}
	baseModel.Widrawal.Amount = createWithdrawals.Amount
	baseModel.Widrawal.FundCode = createWithdrawals.FundCode
	baseModel.Widrawal.AccountUuid = userDtls.AcntUuid
	baseModel.Widrawal.PartnerTransactionId = createWithdrawals.PartnerTransactionId
	response, err := p.TSAClient.SendPostRequest(constants.CreateWithdrawals, &baseModel)
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
