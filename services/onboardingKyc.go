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

func (p *ServiceConfig)ReadAddressProof(readAddressProof *models.)