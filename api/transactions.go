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

func (s *HTTPServer) GetTransactionController(c *fiber.Ctx) error {

	//userId from the bearer token
	userId := c.Locals("userId").(string)
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	fmt.Println(userId)
	baseModel := models.GetTransaction{}

	baseModel.UserId = userId
	//fmt.Println(baseModel.UserId)

	responseCode, data, err := s.MfSrv.GetTransactions(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

// request status
func (s *HTTPServer) RequestStatusController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	baseModel := models.GetTransaction{}

	baseModel.UserId = userId
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	fmt.Println(userId)
	responseCode, data, err := s.MfSrv.RequestStatusCode(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) CurrentInvestedValueController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	baseModel := models.CurrentInvestedValue{}

	baseModel.UserId = userId
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	responseCode, data, err := s.MfSrv.CurrentInvestedValue(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}
func (s *HTTPServer) ReturnsInterestCalculatorController(c *fiber.Ctx) error {
	baseModel := models.ReturnsCalc{}
	customErrors, err := ValidateRequest[models.ReturnsCalc](s, c, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}
	responseCode, data, err := s.MfSrv.ReturnsInterestCalculator(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) RecommendationController(c *fiber.Ctx) error {
	////userId from the bearer token
	//userId := c.Locals("userId").(string)
	//baseModel := models.Recommendation{}
	//
	//baseModel.UserId = userId
	//if userId == "" {
	//	errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	//}
	responseCode, data, err := s.MfSrv.Recommendations()
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}
func (s *HTTPServer) PopularFundsController(c *fiber.Ctx) error {
	////userId from the bearer token
	//userId := c.Locals("userId").(string)
	//baseModel := models.PopularFunds{}

	//baseModel.UserId = userId
	//if userId == "" {
	//	errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	//}
	responseCode, data, err := s.MfSrv.PopularFunds()
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}
func (s *HTTPServer) FundCategoriesController(c *fiber.Ctx) error {

	responseCode, data, err := s.MfSrv.FundCategories()
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}
func (s *HTTPServer) DistinctCategoriesController(c *fiber.Ctx) error {
	responseCode, data, err := s.MfSrv.DistinctFunds()
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) AddToWatchListController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	baseModel := models.AddToWatchList{}
	customErrors, err := ValidateRequest[models.AddToWatchList](s, c, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}
	baseModel.UserId = userId
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	responseCode, data, err := s.MfSrv.AddToWatchList(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) ShowWatchListController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	baseModel := models.ShowWatchList{}
	baseModel.UserId = userId
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	responseCode, data, err := s.MfSrv.ShowWatchList(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

//func (s *HTTPServer) SortedTransactionController(c *fiber.Ctx) error {
//	//userId from the bearer token
//	userId := c.Locals("userId").(string)
//	baseModel := models.UserDtls{}
//	//fmt.Println(userId)
//	baseModel.UserId = userId
//	if userId == "" {
//		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
//	}
//	//fmt.Println(baseModel.UserId)
//	responseCode, data, err := s.MfSrv.SortedTransaction(&baseModel)
//	if err != nil {
//		fmt.Print(err)
//		utils.Log.Error(err)
//		SendResponse(c, responseCode, 0, "processing error", nil, err)
//		return nil
//	}
//	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
//	return nil
//}
