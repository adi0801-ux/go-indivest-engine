package services

import (
	"bytes"
	"encoding/json"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"io"
	"mime/multipart"
	"net/http"
)

func (p *ServiceConfig) StartFullKyc(userDetails *models.StartFullKyc) (int, interface{}, error) {
	baseModel := models.StartFullKycAPI{}
	baseModel.Name = userDetails.Name
	baseModel.Email = userDetails.Email
	baseModel.PhoneNumber = userDetails.PhoneNumber
	baseModel.FullKycRedirectUrl = userDetails.FullKycRedirectUrl

	response, err := p.TSAClient.SendPostRequest(constants.StartFullKyc, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		//	return http.StatusBadRequest, err
	}
	fullKyc := &models.StartFullKycDB{}

	err = p.FullKycRepo.CreateFullKyc(fullKyc)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, nil

}

func (p *ServiceConfig) UploadFile(uploadFile *models.UploadFile) (int, interface{}, error) {
	payload := new(bytes.Buffer)
	writer := multipart.NewWriter(payload)
	file, err := uploadFile.UploadFile.Open()
	defer file.Close()
	filePart, errFile := writer.CreateFormFile("upload", uploadFile.UploadFile.Filename)
	if errFile != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	_, err = io.Copy(filePart, file)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	header := writer.FormDataContentType()
	response, err := p.TSAClient.SendPostFormRequest(constants.UploadFile, payload, header)
	var data models.UploadFileAPI
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	//save to db
	//uploadFiles := models.UploadFileDB{
	//	UserId : uploadFile.UserId,
	//	Url : data.Url,
	//}

	return response.StatusCode, map[string]string{"url": data.Url}, err
}

func (p *ServiceConfig) ReadPanCard(readPanCard *models.ReadPanCard) (int, interface{}, error) {
	baseModel := models.ReadPanCardAPI{}
	baseModel.ImageUrl = readPanCard.ImageUrl
	response, err := p.TSAClient.SendPostRequest(constants.ReadPanCard, &baseModel)
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
	readPanCards := models.ReadPanCardDB{}
	err = p.ReadPanCardRepo.CreateReadPanCardDetails(&readPanCards)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, nil
}

func (p *ServiceConfig) SubmitPanCard(submitPanCard *models.SubmitPanCard) (int, interface{}, error) {
	baseModel := models.SubmitPanCardAPI{}
	baseModel.Name = submitPanCard.Name
	baseModel.FathersName = submitPanCard.FathersName
	baseModel.DateOfBirth = submitPanCard.DateOfBirth
	baseModel.PanNumber = submitPanCard.PanNumber

	response, err := p.TSAClient.SendPostRequest(constants.SubmitPanCard, &baseModel)
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
	return response.StatusCode, nil, nil
}

func (p *ServiceConfig) ReadAddressProof(readAddressProof *models.ReadAddressProof) (int, interface{}, error) {
	baseModel := models.ReadAddressProofAPI{}
	baseModel.UserId = readAddressProof.UserId
	baseModel.AddressProofType = readAddressProof.AddressProofType
	baseModel.ImageUrl = readAddressProof.ImageUrl

	response, err := p.TSAClient.SendPostRequest(constants.ReadAddressProof, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.ReadAddressProofAPIResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	readAddProof := &models.ReadAddressProofDB{
		AadharUid:      readAddressProof.UserId,
		LicenceNumber:  data.LicenceNumber,
		PassportNumber: data.PassportNumber,
		VoterIdNumber:  data.VoterIdNumber,
		Name:           data.Name,
		DateOfBirth:    data.DateOfBirth,
		PinCode:        data.PinCode,
		Address:        data.Address,
		District:       data.District,
		City:           data.City,
		State:          data.State,
		IssueDate:      data.IssueDate,
		ExpiryDate:     data.ExpiryDate,
		FathersName:    data.FathersName,
	}
	err = p.ReadAddressProofRepo.CreateReadAddressProof(readAddProof)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, nil
}

func (p *ServiceConfig) SubmitAddressProof(submitAddressProof *models.SubmitAddressProof) (int, interface{}, error) {
	baseModel := models.SubmitAddressProofAPI{}
	baseModel.AddressProofType = submitAddressProof.AddressProofType
	baseModel.Name = submitAddressProof.Name
	baseModel.ExpiryDate = submitAddressProof.ExpiryDate
	baseModel.DateOfBirth = submitAddressProof.DateOfBirth
	baseModel.IssueDate = submitAddressProof.IssueDate
	baseModel.Address = submitAddressProof.Address
	baseModel.City = submitAddressProof.City
	baseModel.State = submitAddressProof.State
	baseModel.District = submitAddressProof.District
	baseModel.PinCode = submitAddressProof.PinCode
	baseModel.LicenceNumber = submitAddressProof.LicenceNumber
	baseModel.AadharUid = submitAddressProof.AadharUid
	baseModel.PassportNumber = submitAddressProof.PassportNumber
	baseModel.VoterIdNumber = submitAddressProof.VoterIdNumber

	response, err := p.TSAClient.SendPostRequest(constants.SubmitAddressProof, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.SubmitAddressProofAPIResponse
	//converting struct to []bytes
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	return response.StatusCode, nil, nil
}

func (p *ServiceConfig) SubmitInvestorDetails(submitInvestor *models.SubmitInvestorDetails) (int, interface{}, error) {
	baseModel := models.SubmitInvestorDetailsAPI{}
	baseModel.Gender = submitInvestor.Gender
	baseModel.MaritalStatus = submitInvestor.MaritalStatus
	baseModel.OccupationDescription = submitInvestor.OccupationDescription
	baseModel.OccupationCode = submitInvestor.OccupationCode
	baseModel.CitizenshipCode = submitInvestor.CitizenshipCode
	baseModel.CitizenshipCountry = submitInvestor.CitizenshipCountry
	baseModel.ApplicationStatusCode = submitInvestor.ApplicationStatusCode
	baseModel.ApplicationStatusDescription = submitInvestor.ApplicationStatusDescription
	baseModel.AnnualIncome = submitInvestor.AnnualIncome

	response, err := p.TSAClient.SendPostRequest(constants.SubmitInvestorDetails, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		return http.StatusBadRequest, nil, err
	}
	var data models.SubmitInvestorDetailsAPIResponse
	//convert struct to []byte
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		utils.Log.Error(err)
		return response.StatusCode, nil, err
	}
	return response.StatusCode, nil, nil
}
