package services

import (
	"encoding/json"
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
	"net/url"
	"time"
)

// Deposits API
func (p *MFService) GetDeposits(getDeposits *models.GetDeposits) (int, interface{}, error) {
	userDtls, err := p.ShowAccountRepo.ReadAccount(getDeposits.UserId)
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
	onboardingObject, err := p.SavvyRepo.ReadOnboardingObject(createDeposit.UserId)

	userDtls, err := p.ShowAccountRepo.ReadAccount(createDeposit.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		userDtls.AcntUuid = ""
		return http.StatusBadRequest, nil, err
	}
	baseModel := models.CreateDepositAPI{}

	baseModel.Deposit.Amount = createDeposit.Amount
	baseModel.Deposit.FundCode = createDeposit.FundCode
	//baseModel.PaymentRedirectUrl = p.config.PaymentRedirectUrl

	if userDtls.AcntUuid == "" {
		//acntuuid is not present
		baseModel.Deposit.OnboardingUuid = onboardingObject.Uuid
	} else {
		//acntuuid is present
		baseModel.Deposit.AccountUuid = userDtls.AcntUuid
		baseModel.Deposit.OnboardingUuid = onboardingObject.Uuid
	}
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
	if userDtls.AcntUuid == "" {
		userDtls := &models.ShowAccountDB{}
		userDtls.AcntUuid = data.Deposit.Uuid
		userDtls.UserId = createDeposit.UserId
		err = p.ShowAccountRepo.CreateAccount(userDtls)
		if err != nil {
			utils.Log.Error(err)
			return http.StatusBadRequest, nil, err
		}
	}

	createDB := &models.CreateDepositsDb{
		UserId:            createDeposit.UserId,
		FundCode:          data.Deposit.FundCode,
		Amount:            data.Deposit.Amount,
		PaymentStatus:     "payment initiated",
		TransactionStatus: "transaction initiated",
		CreatedAt:         time.Now().UTC(),
	}
	//create update query
	err = p.SavvyRepo.CreateDeposits(createDB)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, map[string]string{"payment_url": data.Url}, nil
}

// if status code was 200 ,
// create an entry in db for deposits against user id , --> fields fund code , amount  , payment status , transaction status , created_at , updated_at , payment_confirmation_time , asset_allocation_time

//return the payment url  {"payment_link":{url}}

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
	withdrawal, err := p.SavvyRepo.ReadWithdrawal(verifyOtp.WithdrawalId)
	fmt.Print(withdrawal.Uuid)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.VerifyWithdrawalOtpAPI{}
	baseModel.Withdrawal.Otp = verifyOtp.Otp
	response, err := p.TSAClient.SendPostRequest(constants.GenerateVerifyWithdrawalOtpUrl(withdrawal.Uuid), &baseModel)
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

	withdrawal.WithdrawalStatus = constants.WithdrawalComplete
	err = p.SavvyRepo.UpdateWithdrawal(withdrawal)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, nil, nil

}
func (p *MFService) CreateWithdrawal(createWithdrawal *models.CreateWithdrawals) (int, interface{}, error) {
	userDtls, err := p.ShowAccountRepo.ReadAccount(createWithdrawal.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.CreateWithdrawalAPI{}
	baseModel.Withdrawal.Amount = createWithdrawal.Amount
	baseModel.Withdrawal.FundCode = createWithdrawal.FundCode
	baseModel.Withdrawal.AccountUuid = userDtls.AcntUuid
	baseModel.Withdrawal.PartnerTransactionId = utils.GeneratePartnerTransactionID()
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
	createWithdrawals := &models.CreateWithdrawalDb{
		UserId:           createWithdrawal.UserId,
		Uuid:             data.Withdrawal.Uuid,
		Amount:           data.Withdrawal.Amount,
		FundCode:         data.Withdrawal.FundCode,
		FundName:         data.Withdrawal.FundName,
		WithdrawalStatus: constants.WithdrawalInitiated,
		WithdrawlId:      utils.GenerateWithdrwalId(),
	}
	err = p.SavvyRepo.CreateWithdrawal(createWithdrawals)
	return response.StatusCode, map[string]string{"withdrawal_id": createWithdrawals.WithdrawlId}, nil
}

//Sip API

func (p *MFService) GetSip(getSip *models.GetSip) (int, interface{}, error) {
	userDtls, err := p.ShowAccountRepo.ReadAccount(getSip.UserId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}
	baseModel := models.GetSipAPI{}
	baseModel.AccountUuid = userDtls.AcntUuid
	params := url.Values{}
	params.Add("account_uuid", userDtls.AcntUuid)
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

	onboardingObject, err := p.SavvyRepo.ReadOnboardingObject(createSip.UserId)

	userDtls, err := p.ShowAccountRepo.ReadAccount(createSip.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		userDtls.AcntUuid = ""
		return http.StatusBadRequest, nil, err
		//	set account uuid as empty
		//	update on completion
		//	case 1 when account id not preset  , call the deposit api and save the account uuid in db of userAccDtls(something like that)
	}
	baseModel := models.CreateSipAPI{}

	baseModel.Sip.Amount = createSip.Amount
	baseModel.Sip.FundCode = createSip.FundCode
	baseModel.Sip.PartnerTransactionId = utils.GeneratePartnerTransactionID()
	baseModel.Sip.StartDate = createSip.StartDate
	baseModel.Sip.EndDate = createSip.EndDate
	baseModel.Sip.Frequency = "monthly"
	baseModel.Sip.MandateRedirectUrl = "p.config.RedirectUrl"
	if userDtls.AcntUuid == "" {
		//acntuuid is not present
		baseModel.Sip.OnboardingUuid = onboardingObject.Uuid
	} else {
		//acntuuid is present
		baseModel.Sip.AccountUuid = userDtls.AcntUuid
	}

	response, err := p.TSAClient.SendPostRequest(constants.CreateSip, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.CreateSipApiResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	createDB := &models.CreateSipDb{
		UserId:        createSip.UserId,
		StartDate:     data.Sip.StartDate,
		EndDate:       data.Sip.EndDate,
		Frequency:     data.Sip.Frequency,
		FundCode:      data.Sip.FundCode,
		Amount:        data.Sip.Amount,
		Uuid:          data.Sip.Uuid,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Now(),
		PaymentStatus: "Payment Pending",
	}
	//create query
	err = p.SavvyRepo.CreateSip(createDB)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, map[string]string{"payments url": data.Url}, err
}

func (p *MFService) RequestStatusCode(requestStatus string) (int, interface{}, error) {
	if requestStatus == "SUCCESS" {
		return http.StatusOK, requestStatus, nil
	} else if requestStatus == "FAILURE" {
		return http.StatusBadRequest, requestStatus, nil
	} else {
		return 0, requestStatus, nil
	}
}

func (p *MFService) GetHoldings(holdings *models.Holding) (int, interface{}, error) {

	params := url.Values{}
	params.Add("amc_code", holdings.AmcCode)
	response, err := p.TSAClient.SendGetRequest(constants.GenerateHoldingsURL(holdings.FundCode), params)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, data, nil
}
