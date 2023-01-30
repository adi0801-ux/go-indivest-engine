package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
)

func (s *HTTPServer) CreateWithdrawlController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	baseModel := models.CreateWithdrawls{}
	baseModel.UserId = userId
	responseCode, data, err := s.MfSrv.CreateWithdrawl(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err)
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}
