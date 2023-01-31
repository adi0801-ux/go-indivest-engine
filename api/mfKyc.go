package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
)

func (s *HTTPServer) CheckIfKycDoneController(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)

	var KYCUser models.CheckKYCUser

	customErrors, err := ValidateRequest[models.CheckKYCUser](s, c, &KYCUser)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	KYCUser.UserId = userID

	responseCode, data, err := s.MfSrv.CheckIfKycDone(&KYCUser)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) StartFullKycController(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)

	var userDetails models.StartFullKyc

	customErrors, err := ValidateRequest[models.StartFullKyc](s, c, &userDetails)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	userDetails.UserId = userID

	responseCode, data, err := s.MfSrv.StartFullKyc(&userDetails)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) AddBankAccountController(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)

	var userDetails models.AddBankAccount

	customErrors, err := ValidateRequest[models.AddBankAccount](s, c, &userDetails)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	userDetails.UserId = userID

	responseCode, data, err := s.MfSrv.AddBankAccount(&userDetails)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) GetOccupationsController(c *fiber.Ctx) error {

	responseCode, data, err := s.MfSrv.GetOccupationStatus()
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) GetGenderCodesController(c *fiber.Ctx) error {

	responseCode, data, err := s.MfSrv.GetGenderCodes()
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) GetMaritalStatusCodesController(c *fiber.Ctx) error {

	responseCode, data, err := s.MfSrv.GetMaritalStatusCodes()
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) GetCountryCodesController(c *fiber.Ctx) error {

	responseCode, data, err := s.MfSrv.GetCountryCodes()
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) GetAnnualIncomeLevelController(c *fiber.Ctx) error {

	responseCode, data, err := s.MfSrv.GetAnnualIncomeLevel()
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) AddPersonalDetailsController(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)

	var userDetails models.AddPersonalDetails

	customErrors, err := ValidateRequest[models.AddPersonalDetails](s, c, &userDetails)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	userDetails.UserId = userID

	responseCode, data, err := s.MfSrv.AddPersonalDetails(&userDetails)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}
	
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) UploadPanCardController(c *fiber.Ctx) error {

	var uploadPan models.UploadPanCard

	//get file
	file, err := c.FormFile("pan_card")
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), err)
		return nil
	}

	uploadPan.PanCard = file
	uploadPan.UserId = c.Locals("userId").(string)

	responseCode, data, err := s.MfSrv.UploadPanCardImage(&uploadPan)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) SubmitPanCardController(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)

	var submitPanCard models.SubmitPanCard

	customErrors, err := ValidateRequest[models.SubmitPanCard](s, c, &submitPanCard)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	submitPanCard.UserId = userID

	responseCode, data, err := s.MfSrv.SubmitPanCard(&submitPanCard)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) UploadAadhaarCardController(c *fiber.Ctx) error {

	var uploadAadhaarCard models.UploadAadhaarCard

	//get file
	file, err := c.FormFile("aadhaar_card")
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), err)
		return nil
	}

	uploadAadhaarCard.AadhaarCard = file
	uploadAadhaarCard.UserId = c.Locals("userId").(string)

	responseCode, data, err := s.MfSrv.UploadAadhaarCardImage(&uploadAadhaarCard)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) SubmitAadhaarCardController(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)

	var submitAadhaarCardImage models.SubmitAadhaarCardImage

	customErrors, err := ValidateRequest[models.SubmitAadhaarCardImage](s, c, &submitAadhaarCardImage)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	submitAadhaarCardImage.UserId = userID

	responseCode, data, err := s.MfSrv.SubmitAadhaarCardImage(&submitAadhaarCardImage)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) SubmitInvestorDetailsController(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)

	var submitInvestorDetails models.SubmitInvestorDetails

	customErrors, err := ValidateRequest[models.SubmitInvestorDetails](s, c, &submitInvestorDetails)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	submitInvestorDetails.UserId = userID

	responseCode, data, err := s.MfSrv.SubmitInvestorDetails(&submitInvestorDetails)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) UploadSignatureController(c *fiber.Ctx) error {

	var uploadSignature models.UploadSignature

	//get file
	file, err := c.FormFile("signature")
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), err)
		return nil
	}

	uploadSignature.Signature = file
	uploadSignature.UserId = c.Locals("userId").(string)

	responseCode, data, err := s.MfSrv.UploadSignature(&uploadSignature)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) UploadSelfieController(c *fiber.Ctx) error {

	var uploadSelfie models.UploadSelfie

	//get file
	file, err := c.FormFile("selfie")
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), err)
		return nil
	}

	uploadSelfie.Selfie = file
	uploadSelfie.UserId = c.Locals("userId").(string)

	responseCode, data, err := s.MfSrv.UploadSelfie(&uploadSelfie)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) StartVideoVerificationController(c *fiber.Ctx) error {
	userID := c.Locals("userId").(string)

	var startVideoVerification models.StartVideoVerification

	customErrors, err := ValidateRequest[models.StartVideoVerification](s, c, &startVideoVerification)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	startVideoVerification.UserId = userID

	responseCode, data, err := s.MfSrv.StartVideoVerification(&startVideoVerification)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) SubmitVideoVerificationController(c *fiber.Ctx) error {

	var submitVideoVerification models.SubmitVideoVerification

	//get file
	file, err := c.FormFile("video")
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), err)
		return nil
	}

	submitVideoVerification.Video = file
	submitVideoVerification.UserId = c.Locals("userId").(string)

	responseCode, data, err := s.MfSrv.SubmitVideoVerification(&submitVideoVerification)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) GenerateKYCContractController(c *fiber.Ctx) error {

	var generateKycContract models.GenerateKycContract

	generateKycContract.UserId = c.Locals("userId").(string)

	responseCode, data, err := s.MfSrv.GenerateKycContract(&generateKycContract)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) ExecuteKYCVerificationController(c *fiber.Ctx) error {

	var executeVerification models.ExecuteVerification

	executeVerification.UserId = c.Locals("userId").(string)

	responseCode, data, err := s.MfSrv.ExecuteVerification(&executeVerification)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}

	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}
