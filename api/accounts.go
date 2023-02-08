package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
)

func (s *HTTPServer) ShowAccountDetailsController(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	baseModel := models.ShowAccount{}
	baseModel.UserId = userId
	responseCode, data, err := s.MfSrv.ShowAccounts(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "proccessing error", nil, err.Error())
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

// Webhooks
func (s *HTTPServer) ConnectWebhooksController(c *fiber.Ctx) error {
	baseModel := models.Webhook{}
	customErrors, err := ValidateRequest[models.Webhook](s, c, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}
	responseCode, data, err := s.MfSrv.ConnectWebhooks(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}
