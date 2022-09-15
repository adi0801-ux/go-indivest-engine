package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"indivest-engine/constants"
	"indivest-engine/models"
	"net/http"
)

func (s *HTTPServer) sandboxBuyMutualFund(c *fiber.Ctx) error {
	var userTransaction models.BuyMutualFund

	customErrors, err := ValidateRequest[models.BuyMutualFund](s, c, &userTransaction)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	userTransaction.UserId = c.Locals("userId").(string)

	err = s.SandboxSrv.BuyMutualFund(&userTransaction)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return nil
	}

	SendSuccessResponse(c, fiber.StatusOK, 1, "SUCCESS", nil)
	return nil
}

func (s *HTTPServer) sandboxRedeemMutualFund(c *fiber.Ctx) error {
	var userTransaction models.RedeemMutualFund

	customErrors, err := ValidateRequest[models.RedeemMutualFund](s, c, &userTransaction)
	if err != nil {
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	userTransaction.UserId = c.Locals("userId").(string)

	err = s.SandboxSrv.RedeemMutualFund(&userTransaction)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return nil
	}

	SendSuccessResponse(c, fiber.StatusOK, 1, "SUCCESS", nil)
	return nil
}

func (s *HTTPServer) sandboxGetHolding(c *fiber.Ctx) error {
	schemeCode := c.Query("scheme_code")
	if schemeCode == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError+" scheme_code"))
	}
	userID := c.Locals("userId").(string)

	holding, err := s.SandboxSrv.GetUserHolding(userID, schemeCode)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return nil
	}

	SendSuccessResponse(c, fiber.StatusOK, 1, "SUCCESS", holding)
	return nil
}

func (s *HTTPServer) sandboxGetAllHolding(c *fiber.Ctx) error {

	userID := c.Locals("userId").(string)

	holding, err := s.SandboxSrv.GetAllUserHoldings(userID)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return nil
	}

	SendSuccessResponse(c, fiber.StatusOK, 1, "SUCCESS", holding)
	return nil
}

func (s *HTTPServer) sandboxGetWallet(c *fiber.Ctx) error {

	userID := c.Locals("userId").(string)

	wallet, err := s.SandboxSrv.GetUserWallet(userID)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return nil
	}

	SendSuccessResponse(c, fiber.StatusOK, 1, "SUCCESS", wallet)
	return nil
}

func (s *HTTPServer) sandboxGetTransactions(c *fiber.Ctx) error {

	schemeCode := c.Query("scheme_code")

	userID := c.Locals("userId").(string)

	var transactions []models.UserMFTransactions
	var err error

	if schemeCode == "" {
		transactions, err = s.SandboxSrv.GetUserAllTransactions(userID)
	} else {
		transactions, err = s.SandboxSrv.GetUserTransactions(userID, schemeCode)
	}
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return nil
	}

	SendSuccessResponse(c, fiber.StatusOK, 1, "SUCCESS", transactions)
	return nil
}
