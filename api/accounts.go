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
