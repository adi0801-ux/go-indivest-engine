package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
	"net/http"
	"strings"
)

func (s *HTTPServer) UserSignUpController(c *fiber.Ctx) error {
	//userId from the bearer token

	baseModel := models.UserSignup{}
	customErrors, err := ValidateRequest[models.UserSignup](s, c, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	responseCode, data, err := s.UserSrv.UserSignUp(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) UserLoginController(c *fiber.Ctx) error {
	//userId from the bearer token

	baseModel := models.UserLogin{}
	customErrors, err := ValidateRequest[models.UserLogin](s, c, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	responseCode, data, err := s.UserSrv.UserLogin(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) VerifyTokenController(c *fiber.Ctx) error {

	clientToken := c.GetReqHeaders()["Authorization"]

	if clientToken == "" {

		errorResponse(c, http.StatusBadRequest, fmt.Errorf("no Authorization header provided"))

		return nil
	}

	extractedToken := strings.Split(clientToken, "Bearer ")

	if len(extractedToken) == 2 {
		clientToken = strings.TrimSpace(extractedToken[1])

	} else {

		errorResponse(c, http.StatusBadRequest, fmt.Errorf("incorrect format of authorization token"))

		return nil
	}

	responseCode, data, err := s.UserSrv.UserVerifyToken(clientToken)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) CreateOnBoardingQuestionsController(c *fiber.Ctx) error {
	//userId from the bearer token
	userId := c.Locals("userId").(string)
	if userId == "" {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError))
	}
	baseModel := models.UserQuestioner{}
	customErrors, err := ValidateRequest[models.UserQuestioner](s, c, &baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendFullErrorResponse(c, http.StatusBadRequest, fmt.Errorf(constants.RequestError), customErrors)
		return nil
	}

	baseModel.UserId = userId
	responseCode, data, err := s.UserSrv.SaveOnBoardingQuestionnaire(&baseModel)
	if err != nil {
		utils.Log.Error(err)
		SendResponse(c, responseCode, 0, "processing error", nil, err.Error())
		return nil
	}
	SendSuccessResponse(c, responseCode, 1, "SUCCESS", data)
	return nil
}
