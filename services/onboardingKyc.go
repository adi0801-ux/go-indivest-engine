package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"io"
	"mime/multipart"
	"net/http"
)

// check if customer kyc exists
func (p *MFService) CheckIfKycDone(KYCUser *models.CheckKYCUser) (int, interface{}, error) {

	//check whether user already registered with us
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(KYCUser.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}
	if onBoardingObject.CreatedAt.String() != constants.StartDateTime {
		utils.Log.Debug("user calling this again with user_id: ", KYCUser.UserId)
		return http.StatusOK, map[string]interface{}{"existing_investor": onBoardingObject.ExistingInvestor, "user_meta": map[string]string{"user_id": onBoardingObject.Uuid}}, nil
	}

	baseModel := models.CheckKYCUserAPI{}

	baseModel.Onboarding.PanNumber = KYCUser.PanNumber
	baseModel.Onboarding.AmcCode = constants.DefaultAMCCode

	response, err := p.TSAClient.SendPostRequest(constants.OnboardingsEndpoint, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.OnBoardingObject
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	//we need to save the user_id against a uuid received for every user
	onboardingDb := &models.OnboardingObjectDB{
		UserId:               KYCUser.UserId,
		Uuid:                 data.Onboarding.Uuid,
		PanNumber:            data.Onboarding.PanNumber,
		ExistingInvestor:     "",
		Name:                 data.Onboarding.Name,
		Email:                data.Onboarding.Email,
		PhoneNumber:          data.Onboarding.PhoneNumber,
		DateOfBirth:          data.Onboarding.DateOfBirth,
		KycStatus:            data.Onboarding.KycStatus,
		PanCardImageUrl:      data.Onboarding.PanCardImageUrl,
		FathersName:          data.Onboarding.FathersName,
		AddressProofImageUrl: data.Onboarding.AddressProofImageUrl,
		AddressProofType:     data.Onboarding.AddressProofType,
		Address:              data.Onboarding.Address,
		City:                 data.Onboarding.City,
		Pincode:              data.Onboarding.Pincode,
		SignatureImageUrl:    data.Onboarding.SignatureImageUrl,
		SelfieImageUrl:       data.Onboarding.SelfieImageUrl,
		CancelledChequeUrl:   data.Onboarding.CancelledChequeUrl,
		VideoUrl:             data.Onboarding.VideoUrl,
		AnnualIncome:         data.Onboarding.AnnualIncome,
		Gender:               data.Onboarding.Gender,
		Occupation:           data.Onboarding.Occupation,
		MaritalStatus:        data.Onboarding.MaritalStatus,
	}

	if data.Onboarding.ExistingInvestor == true {
		onboardingDb.ExistingInvestor = "1"
	} else {
		onboardingDb.ExistingInvestor = "0"
	}

	//save to db
	err = p.SavvyRepo.CreateOnboardingObject(onboardingDb)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, map[string]interface{}{"existing_investor": onboardingDb.ExistingInvestor, "user_meta": map[string]string{"user_id": onboardingDb.Uuid}}, nil

}

// helper apis
func (p *MFService) GetOccupationStatus() (int, interface{}, error) {
	data, err := p.SavvyRepo.ReadAllOccupationStatus()
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, data, nil
}

func (p *MFService) GetGenderCodes() (int, interface{}, error) {
	data, err := p.SavvyRepo.ReadAllGenderCodes()
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, data, nil
}

func (p *MFService) GetMaritalStatusCodes() (int, interface{}, error) {
	data, err := p.SavvyRepo.ReadAllMaritalStatusCodes()
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, data, nil
}

func (p *MFService) GetCountryCodes() (int, interface{}, error) {
	data, err := p.SavvyRepo.ReadAllCountryCodes()
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, data, nil
}

func (p *MFService) GetAnnualIncomeLevel() (int, interface{}, error) {
	data, err := p.SavvyRepo.ReadAllAnnualIncomeLevel()
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return http.StatusOK, data, nil
}

// StartFullKyc step 1 - existing investor

// step 1 - Add bank details
func (p *MFService) AddBankAccount(userDetails *models.AddBankAccount) (int, interface{}, error) {

	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(userDetails.UserId)

	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}
	//if onBoardingObject.ExistingInvestor == "0" {
	//	return http.StatusBadRequest, nil, fmt.Errorf("user kyc required")
	//}

	baseModel := models.AddBankAccountAPI{}
	baseModel.Onboarding.AccountNumber = userDetails.AccountNumber
	baseModel.Onboarding.IfscCode = userDetails.IFSC

	response, err := p.TSAClient.SendPostRequest(constants.GenerateAddBankURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.AddBankAccountAPIResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	//saving onboarding fields
	onboardingDb := &models.OnboardingObjectDB{
		Uuid:                 data.Onboarding.Uuid,
		PanNumber:            data.Onboarding.PanNumber,
		ExistingInvestor:     "",
		Name:                 data.Onboarding.Name,
		Email:                data.Onboarding.Email,
		PhoneNumber:          data.Onboarding.PhoneNumber,
		DateOfBirth:          data.Onboarding.DateOfBirth,
		KycStatus:            data.Onboarding.KycStatus,
		PanCardImageUrl:      data.Onboarding.PanCardImageUrl,
		FathersName:          data.Onboarding.FathersName,
		AddressProofImageUrl: data.Onboarding.AddressProofImageUrl,
		AddressProofType:     data.Onboarding.AddressProofType,
		Address:              data.Onboarding.Address,
		City:                 data.Onboarding.City,
		Pincode:              data.Onboarding.Pincode,
		SignatureImageUrl:    data.Onboarding.SignatureImageUrl,
		SelfieImageUrl:       data.Onboarding.SelfieImageUrl,
		CancelledChequeUrl:   data.Onboarding.CancelledChequeUrl,
		VideoUrl:             data.Onboarding.VideoUrl,
		AnnualIncome:         data.Onboarding.AnnualIncome,
		Gender:               data.Onboarding.Gender,
		Occupation:           data.Onboarding.Occupation,
		MaritalStatus:        data.Onboarding.MaritalStatus,
	}
	err = p.SavvyRepo.UpdateOrCreateOnboardingObject(onboardingDb)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	//save information
	bankAccount := &models.BankAccountDB{
		UserId:        userDetails.UserId,
		AccountNumber: data.BankAccount.AccountNumber,
		BankName:      data.BankAccount.BankName,
		BranchName:    data.BankAccount.BranchName,
		BankCity:      data.BankAccount.BankCity,
		IfscCode:      data.BankAccount.IfscCode,
	}

	err = p.SavvyRepo.CreateUserBank(bankAccount)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil
}

// step 2
func (p *MFService) AddPersonalDetails(userDetails *models.AddPersonalDetails) (int, interface{}, error) {

	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(userDetails.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	if onBoardingObject.ExistingInvestor == "0" {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, fmt.Errorf("user kyc required")
	}

	baseModel := models.AddPersonalDetailsAPI{}
	baseModel.Onboarding.Address = userDetails.Address
	baseModel.Onboarding.City = userDetails.City
	baseModel.Onboarding.Pincode = userDetails.Pincode
	baseModel.Onboarding.DateOfBirth = userDetails.DateOfBirth
	baseModel.Onboarding.Occupation = userDetails.OccupationCode
	baseModel.Onboarding.Fatca.FatcaBirthCountryCode = userDetails.BirthCountryCode
	baseModel.Onboarding.Fatca.FatcaCitizenshipCountryCode = userDetails.CitizenshipCountryCode
	baseModel.Onboarding.Fatca.FatcaTaxCountryCode = userDetails.TaxCountryCode
	baseModel.Onboarding.Fatca.FatcaPlaceOfBirth = userDetails.PlaceOfBirth
	baseModel.Onboarding.Fatca.FatcaAddressType = userDetails.AddressType
	baseModel.Onboarding.Fatca.FatcaOccupation = userDetails.Occupation
	baseModel.Onboarding.Fatca.FatcaGrossIncome = userDetails.GrossIncome
	baseModel.Onboarding.Fatca.FatcaSourceWealth = userDetails.SourceWealth
	response, err := p.TSAClient.SendPutRequest(constants.GenerateAddPersonalDetailsURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	var data models.OnBoardingObject
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	//we need to save the user_id against a uuid received for every user
	onboardingDb := &models.OnboardingObjectDB{
		UserId:               userDetails.UserId,
		Uuid:                 data.Onboarding.Uuid,
		PanNumber:            data.Onboarding.PanNumber,
		ExistingInvestor:     "",
		Name:                 data.Onboarding.Name,
		Email:                data.Onboarding.Email,
		PhoneNumber:          data.Onboarding.PhoneNumber,
		DateOfBirth:          data.Onboarding.DateOfBirth,
		KycStatus:            data.Onboarding.KycStatus,
		PanCardImageUrl:      data.Onboarding.PanCardImageUrl,
		FathersName:          data.Onboarding.FathersName,
		AddressProofImageUrl: data.Onboarding.AddressProofImageUrl,
		AddressProofType:     data.Onboarding.AddressProofType,
		Address:              data.Onboarding.Address,
		City:                 data.Onboarding.City,
		Pincode:              data.Onboarding.Pincode,
		SignatureImageUrl:    data.Onboarding.SignatureImageUrl,
		SelfieImageUrl:       data.Onboarding.SelfieImageUrl,
		CancelledChequeUrl:   data.Onboarding.CancelledChequeUrl,
		VideoUrl:             data.Onboarding.VideoUrl,
		AnnualIncome:         data.Onboarding.AnnualIncome,
		Gender:               data.Onboarding.Gender,
		Occupation:           data.Onboarding.Occupation,
		MaritalStatus:        data.Onboarding.MaritalStatus,
	}

	//create update query
	err = p.SavvyRepo.UpdateOrCreateOnboardingObject(onboardingDb)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil
}

//more steps req

// Case 2 No KYC done

// StartFullKyc step 1 - start with full kyc
func (p *MFService) StartFullKyc(userDetails *models.StartFullKyc) (int, interface{}, error) {

	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(userDetails.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	if onBoardingObject.ExistingInvestor == "1" {
		return http.StatusBadRequest, nil, fmt.Errorf("user already an investor")
	}
	baseModel := models.StartFullKycAPI{}
	baseModel.Onboarding.Name = userDetails.Name
	baseModel.Onboarding.Email = userDetails.Email
	baseModel.Onboarding.PhoneNumber = userDetails.PhoneNumber
	baseModel.Onboarding.FullKycRedirectUrl = constants.RedirectURLAfterKyc

	response, err := p.TSAClient.SendPostRequest(constants.GenerateFullKycURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.OnBoardingObject
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	//we need to save the user_id against a uuid received for every user
	onboardingDb := &models.OnboardingObjectDB{
		UserId:               userDetails.UserId,
		Uuid:                 data.Onboarding.Uuid,
		PanNumber:            data.Onboarding.PanNumber,
		ExistingInvestor:     "",
		Name:                 data.Onboarding.Name,
		Email:                data.Onboarding.Email,
		PhoneNumber:          data.Onboarding.PhoneNumber,
		DateOfBirth:          data.Onboarding.DateOfBirth,
		KycStatus:            data.Onboarding.KycStatus,
		PanCardImageUrl:      data.Onboarding.PanCardImageUrl,
		FathersName:          data.Onboarding.FathersName,
		AddressProofImageUrl: data.Onboarding.AddressProofImageUrl,
		AddressProofType:     data.Onboarding.AddressProofType,
		Address:              data.Onboarding.Address,
		City:                 data.Onboarding.City,
		Pincode:              data.Onboarding.Pincode,
		SignatureImageUrl:    data.Onboarding.SignatureImageUrl,
		SelfieImageUrl:       data.Onboarding.SelfieImageUrl,
		CancelledChequeUrl:   data.Onboarding.CancelledChequeUrl,
		VideoUrl:             data.Onboarding.VideoUrl,
		AnnualIncome:         data.Onboarding.AnnualIncome,
		Gender:               data.Onboarding.Gender,
		Occupation:           data.Onboarding.Occupation,
		MaritalStatus:        data.Onboarding.MaritalStatus,
	}

	//create update query
	err = p.SavvyRepo.UpdateOrCreateOnboardingObject(onboardingDb)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, nil, nil
}

// step 2 UploadFile
func (p *MFService) UploadFile(uploadFile models.UploadFile) (models.UploadFileAPI, error) {

	payload := new(bytes.Buffer)
	writer := multipart.NewWriter(payload)

	file, err := uploadFile.UploadFile.Open()
	if err != nil {
		utils.Log.Error(err)
		return models.UploadFileAPI{}, err
	}

	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			utils.Log.Error(err)
		}
	}(file)

	filePart, errFile := writer.CreateFormFile("upload", uploadFile.UploadFile.Filename)
	if errFile != nil {
		utils.Log.Error(err)
		return models.UploadFileAPI{}, err
	}

	_, err = io.Copy(filePart, file)
	if err != nil {
		utils.Log.Error(err)
		return models.UploadFileAPI{}, err
	}

	header := writer.FormDataContentType()
	err = writer.Close()
	if err != nil {
		utils.Log.Error(err)
		return models.UploadFileAPI{}, err
	}

	response, err := p.TSAClient.SendPostFormRequest(constants.GenerateUploadFileURL(uploadFile.UUID), payload, header)
	if err != nil {
		utils.Log.Error(err)
		return models.UploadFileAPI{}, err
	}
	var data models.UploadFileAPI
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return models.UploadFileAPI{}, err
	}

	return data, nil
}

// step 3 UploadPanCardImage
func (p *MFService) UploadPanCardImage(uploadPan *models.UploadPanCard) (int, interface{}, error) {
	//get uuid and status
	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(uploadPan.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	if onBoardingObject.ExistingInvestor == "1" {
		return http.StatusBadRequest, nil, fmt.Errorf("user already an investor")
	}

	//upload the pan card
	uploadFile := models.UploadFile{
		UUID:       onBoardingObject.Uuid,
		UploadFile: uploadPan.PanCard,
	}

	uploadObject, err := p.UploadFile(uploadFile)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.ReadPanCardAPI{}
	baseModel.Onboarding.ImageUrls = []string{uploadObject.File}
	response, err := p.TSAClient.SendPostRequest(constants.GenerateReadPanCardURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.ReadPanCardAPIResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil
}

// step 4 SubmitPanCard
func (p *MFService) SubmitPanCard(submitPanCard *models.SubmitPanCard) (int, interface{}, error) {
	//get uuid and status
	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(submitPanCard.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	if onBoardingObject.ExistingInvestor == "1" {
		return http.StatusBadRequest, nil, fmt.Errorf("user already an investor")
	}

	baseModel := models.SubmitPanCardAPI{}
	baseModel.Onboarding.Name = submitPanCard.Name
	baseModel.Onboarding.FathersName = submitPanCard.FathersName
	baseModel.Onboarding.DateOfBirth = submitPanCard.DateOfBirth
	baseModel.Onboarding.PanNumber = submitPanCard.PanNumber

	response, err := p.TSAClient.SendPostRequest(constants.GenerateSubmitPanCardURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.SubmitPanCardAPIResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	//we need to save the user_id against a uuid received for every user
	onboardingDb := &models.OnboardingObjectDB{
		UserId:               submitPanCard.UserId,
		Uuid:                 data.Onboarding.Uuid,
		PanNumber:            data.Onboarding.PanNumber,
		ExistingInvestor:     "",
		Name:                 data.Onboarding.Name,
		Email:                data.Onboarding.Email,
		PhoneNumber:          data.Onboarding.PhoneNumber,
		DateOfBirth:          data.Onboarding.DateOfBirth,
		KycStatus:            data.Onboarding.KycStatus,
		PanCardImageUrl:      data.Onboarding.PanCardImageUrl,
		FathersName:          data.Onboarding.FathersName,
		AddressProofImageUrl: data.Onboarding.AddressProofImageUrl,
		AddressProofType:     data.Onboarding.AddressProofType,
		Address:              data.Onboarding.Address,
		City:                 data.Onboarding.City,
		Pincode:              data.Onboarding.Pincode,
		SignatureImageUrl:    data.Onboarding.SignatureImageUrl,
		SelfieImageUrl:       data.Onboarding.SelfieImageUrl,
		CancelledChequeUrl:   data.Onboarding.CancelledChequeUrl,
		VideoUrl:             data.Onboarding.VideoUrl,
		AnnualIncome:         data.Onboarding.AnnualIncome,
		Gender:               data.Onboarding.Gender,
		Occupation:           data.Onboarding.Occupation,
		MaritalStatus:        data.Onboarding.MaritalStatus,
	}

	//create update query
	err = p.SavvyRepo.UpdateOrCreateOnboardingObject(onboardingDb)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, data, nil
}

// step 5 UploadAadhaarCardImage
func (p *MFService) UploadAadhaarCardImage(uploadAadhaar *models.UploadAadhaarCard) (int, interface{}, error) {
	//get uuid and status
	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(uploadAadhaar.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	if onBoardingObject.ExistingInvestor == "1" {
		return http.StatusBadRequest, nil, fmt.Errorf("user already an investor")
	}

	//upload the pan card
	uploadFileAddharFront := models.UploadFile{
		UUID:       onBoardingObject.Uuid,
		UploadFile: uploadAadhaar.AadhaarCardFront,
	}

	uploadObjectAddharFront, err := p.UploadFile(uploadFileAddharFront)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	uploadFileAddharBack := models.UploadFile{
		UUID:       onBoardingObject.Uuid,
		UploadFile: uploadAadhaar.AadhaarCardBack,
	}

	uploadObjectAddharBack, err := p.UploadFile(uploadFileAddharBack)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.UploadAadhaarCardAPI{}
	baseModel.Onboarding.AddressProofType = "aadhaar"
	baseModel.Onboarding.ImageUrls = []string{uploadObjectAddharFront.File, uploadObjectAddharBack.File}

	response, err := p.TSAClient.SendPostRequest(constants.GenerateReadAadharCardURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.UploadAadhaarCardAPIResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	responseData := &models.UploadAadhaarCardResponse{
		AadhaarUid:  data.AadhaarUid,
		AadhaarVid:  data.AadhaarVid,
		Name:        data.Name,
		DateOfBirth: data.DateOfBirth,
		Pincode:     data.Pincode,
		Address:     data.Address,
		District:    data.District,
		City:        data.City,
		State:       data.State,
		Gender:      data.Gender,
	}

	return response.StatusCode, responseData, nil
}

// step 6 SubmitAadhaarCardImage
func (p *MFService) SubmitAadhaarCardImage(submitAadhaarCardImage *models.SubmitAadhaarCardImage) (int, interface{}, error) {

	//get uuid and status
	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(submitAadhaarCardImage.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	if onBoardingObject.ExistingInvestor == "1" {
		return http.StatusBadRequest, nil, fmt.Errorf("user already an investor")
	}

	baseModel := models.SubmitAadhaarCardImageAPI{}
	baseModel.Onboarding.AddressProofType = "aadhaar"
	baseModel.Onboarding.Name = submitAadhaarCardImage.Name
	baseModel.Onboarding.DateOfBirth = submitAadhaarCardImage.DateOfBirth
	baseModel.Onboarding.Address = submitAadhaarCardImage.Address
	baseModel.Onboarding.City = submitAadhaarCardImage.City
	baseModel.Onboarding.State = submitAadhaarCardImage.State
	baseModel.Onboarding.District = submitAadhaarCardImage.District
	baseModel.Onboarding.Pincode = submitAadhaarCardImage.PinCode
	baseModel.Onboarding.AadhaarUid = submitAadhaarCardImage.AadharUid

	response, err := p.TSAClient.SendPostRequest(constants.GenerateSubmitAadharCardURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.SubmitAadhaarCardImageAPIResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, nil
}

// step 7 SubmitInvestorDetails
func (p *MFService) SubmitInvestorDetails(submitInvestorDetails *models.SubmitInvestorDetails) (int, interface{}, error) {
	//get uuid and status
	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(submitInvestorDetails.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	if onBoardingObject.ExistingInvestor == "1" {
		return http.StatusBadRequest, nil, fmt.Errorf("user already an investor")
	}

	baseModel := models.SubmitInvestorDetailsAPI{}
	baseModel.Onboarding.Gender = submitInvestorDetails.Gender
	baseModel.Onboarding.MaritalStatus = submitInvestorDetails.MaritalStatus
	baseModel.Onboarding.OccupationDescription = submitInvestorDetails.OccupationDescription
	baseModel.Onboarding.OccupationCode = submitInvestorDetails.OccupationCode
	baseModel.Onboarding.CitizenshipCode = submitInvestorDetails.CitizenshipCode
	baseModel.Onboarding.CitizenshipCountry = constants.India
	baseModel.Onboarding.ApplicationStatusCode = submitInvestorDetails.ApplicationStatusCode
	baseModel.Onboarding.ApplicationStatusDescription = submitInvestorDetails.ApplicationStatusDescription
	baseModel.Onboarding.AnnualIncome = submitInvestorDetails.AnnualIncome

	response, err := p.TSAClient.SendPostRequest(constants.GenerateInvestorDetailsURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.SubmitInvestorDetailsAPIResponse
	//convert struct to []byte
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, nil
}

// step 8 UploadSignature
func (p *MFService) UploadSignature(uploadSignature *models.UploadSignature) (int, interface{}, error) {
	//get uuid and status
	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(uploadSignature.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	if onBoardingObject.ExistingInvestor == "1" {
		return http.StatusBadRequest, nil, fmt.Errorf("user already an investor")
	}

	//upload the pan card
	uploadFile := models.UploadFile{
		UUID:       onBoardingObject.Uuid,
		UploadFile: uploadSignature.Signature,
	}

	uploadObject, err := p.UploadFile(uploadFile)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.UploadSignatureAPI{}
	baseModel.Onboarding.ImageUrls = []string{uploadObject.File}

	response, err := p.TSAClient.SendPostRequest(constants.GenerateUploadSignatureURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.UploadSignatureAPIResponse
	//convert struct to []byte
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, err
}

// step 9 UploadSelfie
func (p *MFService) UploadSelfie(uploadSelfie *models.UploadSelfie) (int, interface{}, error) {

	//get uuid and status
	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(uploadSelfie.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	if onBoardingObject.ExistingInvestor == "1" {
		return http.StatusBadRequest, nil, fmt.Errorf("user already an investor")
	}

	//upload the pan card
	uploadFile := models.UploadFile{
		UUID:       onBoardingObject.Uuid,
		UploadFile: uploadSelfie.Selfie,
	}

	uploadObject, err := p.UploadFile(uploadFile)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.UploadSelfieAPI{}
	baseModel.Onboarding.ImageUrls = []string{uploadObject.File}

	response, err := p.TSAClient.SendPostRequest(constants.GenerateUploadSelfieURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.UploadSignatureAPIResponse
	//convert struct to []byte
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, err
}

// step 10 StartVideoVerification
func (p *MFService) StartVideoVerification(startVideoVerification *models.StartVideoVerification) (int, interface{}, error) {

	//get uuid and status
	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(startVideoVerification.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	if onBoardingObject.ExistingInvestor == "1" {
		return http.StatusBadRequest, nil, fmt.Errorf("user already an investor")
	}

	var baseModel interface{}
	response, err := p.TSAClient.SendPostRequest(constants.GenerateStartVideoVerificationURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.StartVideoVerificationAPIResponse
	//convert struct to []byte
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	startVideo := &models.StartVideoVerificationDB{
		TransactionId: data.TransactionId,
		RandomNumber:  data.RandomNumber,
		UserId:        startVideoVerification.UserId,
	}
	err = p.SavvyRepo.CreateVideoVerification(startVideo)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, map[string]string{"random_number": data.RandomNumber}, err
}

// step 11 SubmitVideoVerification
func (p *MFService) SubmitVideoVerification(submitVideoVerification *models.SubmitVideoVerification) (int, interface{}, error) {

	//get uuid and status
	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(submitVideoVerification.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	if onBoardingObject.ExistingInvestor == "1" {
		return http.StatusBadRequest, nil, fmt.Errorf("user already an investor")
	}

	//upload the pan card
	uploadFile := models.UploadFile{
		UUID:       onBoardingObject.Uuid,
		UploadFile: submitVideoVerification.Video,
	}

	uploadObject, err := p.UploadFile(uploadFile)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	startVideo, err := p.SavvyRepo.ReadVideoVerification(submitVideoVerification.UserId)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	baseModel := models.SubmitVideoVerificationAPI{}
	baseModel.Onboarding.TransactionId = startVideo.TransactionId
	baseModel.Onboarding.VideoUrl = uploadObject.File

	response, err := p.TSAClient.SendPostRequest(constants.GenerateSubmitVideoVerificationURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.SubmitVideoVerificationAPIResponse
	//convert struct to []byte
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, err
}

// step 12 GenerateKycContract
func (p *MFService) GenerateKycContract(generateKycContract *models.GenerateKycContract) (int, interface{}, error) {

	//get uuid and status
	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(generateKycContract.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	if onBoardingObject.ExistingInvestor == "1" {
		return http.StatusBadRequest, nil, fmt.Errorf("user already an investor")
	}

	var baseModel interface{}
	response, err := p.TSAClient.SendPostRequest(constants.GenerateKYCContractURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.GenerateKycContractAPIResponse
	//convert struct to []byte
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}

	return response.StatusCode, map[string]string{"esign_url": data.Url}, err
}

// step 13 ExecuteVerification
func (p *MFService) ExecuteVerification(executeVerification *models.ExecuteVerification) (int, interface{}, error) {
	//get uuid and status
	//check the exsisting investor status
	onBoardingObject, err := p.SavvyRepo.ReadOnboardingObject(executeVerification.UserId)
	if err != nil && err.Error() != constants.UserNotFound {
		return http.StatusBadRequest, nil, err
	}

	if onBoardingObject.ExistingInvestor == "1" {
		return http.StatusBadRequest, nil, fmt.Errorf("user already an investor")
	}

	var baseModel interface{}
	response, err := p.TSAClient.SendPostRequest(constants.GenerateKYCContractVerifyURL(onBoardingObject.Uuid), &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.ExecuteVerificationAPIResponse
	//convert struct to []byte
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, data, err
}
