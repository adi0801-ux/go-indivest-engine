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

//GetDeposits API
func (s *HTTPServer) GetDepositsController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	baseModel := models.GetDeposits{}
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

//CreateDeposits API

//Withdrawal API
func (s *HTTPServer) CreateWithdrawalController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	baseModel := models.CreateWithdrawals{}
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
