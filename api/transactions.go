package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
)

//DEPOSITS

// GetDeposits API
func (s *HTTPServer) GetDepositsController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	baseModel := models.GetDeposits{}
	customErrors, err := ValidateRequest[models.GetDeposits](s, c, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}
	baseModel.UserId = userId
	responseCode, data, err := s.MfSrv.GetDeposits(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}
	c.Redirect(s.config.RedirectUrl, 302)
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

//ShowDeposits API
//func (s *HTTPServer) ShowDepositsController(c *fiber.Ctx) error {
//
//}

// CreateDeposits API
func (s *HTTPServer) CreateDepositsController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	baseModel := models.CreateDeposit{}
	customErrors, err := ValidateRequest[models.CreateDeposit](s, c, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}
	baseModel.UserId = userId
	responseCode, data, err := s.MfSrv.CreateDeposit(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

// Withdrawal API
func (s *HTTPServer) CreateWithdrawalController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	baseModel := models.CreateWithdrawals{}
	customErrors, err := ValidateRequest[models.CreateWithdrawals](s, c, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}
	baseModel.UserId = userId
	responseCode, data, err := s.MfSrv.CreateWithdrawal(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}
func (s *HTTPServer) VerifyWithdrawalOtpController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	baseModel := models.VerifyWithdrawalOtp{}
	customErrors, err := ValidateRequest[models.VerifyWithdrawalOtp](s, c, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	baseModel.UserId = userId

	responseCode, data, err := s.MfSrv.VerifyWithdrawalOtp(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

//Systematic Investment Plan Controllers

// Create Sip
func (s *HTTPServer) CreateSipController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	baseModel := models.CreateSip{}
	customErrors, err := ValidateRequest[models.CreateSip](s, c, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}
	baseModel.UserId = userId
	responseCode, data, err := s.MfSrv.CreateSip(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

// get sip
func (s *HTTPServer) GetSipController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	baseModel := models.GetSip{}
	customErrors, err := ValidateRequest[models.GetSip](s, c, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}
	baseModel.UserId = userId
	responseCode, data, err := s.MfSrv.GetSip(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) GetHoldingsController(c *fiber.Ctx) error {
	baseModel := models.Holding{}
	customErrors, err := ValidateRequest[models.Holding](s, c, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}
	responseCode, data, err := s.MfSrv.GetHoldings(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

//
//request status
//func (s *HTTPServer)RequestStatus(c *fiber.Ctx)error{
//	responseCode, data, err:= s.MfSrv.RequestStatusCode()
//	if err != nil {
//		utils.Log.Error(err)
//		SendResponse(c, responseCode, 0, "processing error", nil, err)
//		return nil
//	}
//	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
//	return nil
//}
//}