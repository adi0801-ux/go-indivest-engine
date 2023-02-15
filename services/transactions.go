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
	userDtls, err := p.SavvyRepo.ReadAccount(getDeposits.UserId)
	baseModel := models.GetDepositsAPI{}
	baseModel.AccountUuid = userDtls.AcntUuid
	params := url.Values{}
	params.Add("account_uuid", userDtls.AcntUuid)
	fmt.Println(params)
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
	return response.StatusCode, data, err
}

func (p *MFService) ShowDeposits() (int, interface{}, error) {
	return 0, nil, nil
}

func (p *MFService) CreateDeposit(createDeposit *models.CreateDeposit) (int, interface{}, error) {
	onboardingObject, err := p.SavvyRepo.ReadOnboardingObject(createDeposit.UserId)

	userDtls, err := p.SavvyRepo.ReadAccount(createDeposit.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		userDtls.UserId = ""
		userDtls.AcntUuid = ""
		userDtls.AmcId = ""
		return http.StatusBadRequest, nil, err
	}
	baseModel := models.CreateDepositAPI{}

	baseModel.Deposit.Amount = createDeposit.Amount
	baseModel.Deposit.FundCode = createDeposit.FundCode
	//baseModel.PaymentRedirectUrl = p.config.PaymentRedirectUrl

	if userDtls.AcntUuid == "" && userDtls.AmcId == "" {
		//userId and amcId not present
		baseModel.Deposit.OnboardingUuid = onboardingObject.Uuid
	} else {
		//acntUuid is present
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
		err = p.SavvyRepo.CreateAccount(userDtls)
		if err != nil {
			utils.Log.Error(err)
			return http.StatusBadRequest, nil, err
		}
	}

	createDB := &models.CreateDepositsDb{
		Uuid:              data.Deposit.Uuid,
		UserId:            createDeposit.UserId,
		FundCode:          data.Deposit.FundCode,
		NAV:               data.Deposit.NAV,
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
	if err != nil {
		utils.Log.Error(err)
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
	fmt.Print(createWithdrawal.UserId)
	userDtls, err := p.SavvyRepo.ReadAccount(createWithdrawal.UserId)
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
		WithdrawalId:     utils.GenerateWithdrawalId(),
	}
	err = p.SavvyRepo.CreateWithdrawal(createWithdrawals)
	return response.StatusCode, map[string]string{"withdrawal_id": createWithdrawals.WithdrawalId}, nil
}

//Sip API

func (p *MFService) GetSip(getSip *models.GetSip) (int, interface{}, error) {
	userDtls, err := p.SavvyRepo.ReadAccount(getSip.UserId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

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
	return response.StatusCode, data, err
}

func (p *MFService) ShowSip() (int, interface{}, error) {
	return http.StatusOK, nil, nil
}

func (p *MFService) CreateSip(createSip *models.CreateSip) (int, interface{}, error) {

	onboardingObject, err := p.SavvyRepo.ReadOnboardingObject(createSip.UserId)

	userDtls, err := p.SavvyRepo.ReadAccount(createSip.UserId)
	if err != nil && err.Error() == constants.UserNotFound {
		utils.Log.Info(err)
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

func (p *MFService) RequestStatusCode(rqstStatus *models.GetTransaction) (int, interface{}, error) {
	depositDtls, err := p.SavvyRepo.ReadDeposits(rqstStatus.UserId)
	sipDtls, err := p.SavvyRepo.ReadSip(rqstStatus.UserId)
	withdrawDtls, err := p.SavvyRepo.ReadWithdrawalAll(rqstStatus.UserId)
	if err != nil {
		utils.Log.Info(err)
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, map[string]interface{}{"deposit_status": depositDtls.TransactionStatus, "sip_status": sipDtls.SipStatus, "withdrawal_status": withdrawDtls.WithdrawalStatus}, err
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

func (p *MFService) GetTransactions(transDtls *models.GetTransaction) (int, interface{}, error) {
	withdrawals, err := p.SavvyRepo.ReadWithdrawalAll(transDtls.UserId)
	deposits, err := p.SavvyRepo.ReadDeposits(transDtls.UserId)
	sips, err := p.SavvyRepo.ReadSip(transDtls.UserId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	return http.StatusOK, map[string]interface{}{"sip_details": sips, "withdrawl_details": withdrawals, "deposits": deposits}, nil
}

func (p *MFService) CurrentInvestedValue(currentValue *models.CurrentInvestedValue) (int, interface{}, error) {
	fundDtls, err := p.SavvyRepo.ReadFundDetails(currentValue.FundCode)
	depoDtls, err := p.SavvyRepo.ReadDeposits(currentValue.UserId)
	if err != nil {
		utils.Log.Info(err)
	}
	var units = depoDtls.Amount / depoDtls.NAV
	//correct logic
	//there must be 2 navs. NAV1 at the time of purchasae,
	//						NAV2 at the time of calculating currentValue
	currentVal := units * float64(fundDtls.NAV)
	return http.StatusOK, map[string]interface{}{"current_invested_value": currentVal}, nil
}

//
//// datewise sorting for transaction
//func (p *MFService) DatewiseDeposit(userDtls *models.UserDtls) (int, interface{}, error) {
//	depoDtls, err := p.SavvyRepo.ReadDeposits(userDtls.UserId)
//	withDtls, err := p.SavvyRepo.ReadWithdrawalAll(userDtls.UserId)
//	sipDtls, err := p.SavvyRepo.ReadSip(userDtls.UserId)
//
//}
