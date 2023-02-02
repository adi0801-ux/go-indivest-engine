package services

import (
	"encoding/json"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
	"net/url"
)

//Deposits API
func (p *MFService) GetDeposits(getDeposits *models.GetDeposits) (int, interface{}, error) {
	userDtls, err := p.ShowAccountRepo.ReadShowAccount(getDeposits.UserId)
	baseModel := models.GetDepositsAPI{}
	baseModel.AccountUuid = userDtls.AcntUuid
	params := url.Values{}
	params.Add("account_uuid", userDtls.AcntUuid)
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

func (p *MFService) ShowDeposits() (int, interface{}, error) {
	return 0, nil, nil
}

func (p *MFService) CreateDeposit(createDeposit *models.CreateDeposit) (int, interface{}, error) {
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

func (p *MFService) CreateBasketOfDeposit(createBasketOfDeposit *models.CreateBasketOfDeposits) (int, interface{}, error) {
	baseModel := models.CreateBasketOfDepositsAPI{}
	baseModel.PaymentRedirectUrl = createBasketOfDeposit.PaymentRedirectUrl
	baseModel.AccountUuid = createBasketOfDeposit.AccountUuid
	baseModel.OnBoardingUuid = createBasketOfDeposit.OnBoardingUuid
	baseModel.PartnerTransactionId = createBasketOfDeposit.PartnerTransactionId
	baseModel.DepositsParts = createBasketOfDeposit.DepositsParts

	response, err := p.TSAClient.SendPostRequest(constants.CreateBasketOfDeposits, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.CreateBasketOfDepositsAPIResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, err
}

//Withdrawal API

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

//Sip API

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
