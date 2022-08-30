package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"indivest-engine/constants"
	"indivest-engine/models"
	"net/http"
)


func (s *HTTPServer)basicDetailsLanguageController(c *fiber.Ctx) error {
	//load basic details from request

	var basicDetails models.UserBasicDetailsLanguage

	customErrors , err := ValidateRequest[models.UserBasicDetailsLanguage](s, c , &basicDetails)
	if err!=nil{
		SendFullErrorResponse(c , http.StatusBadRequest , fmt.Errorf(constants.RequestError),customErrors)
		return nil
	}

	//basic details can be used here
	err = s.Srv.AddLanguage(&basicDetails)
	if err != nil {
		errorResponse(c , http.StatusBadRequest , err)
		return nil
	}

	SendSuccessResponse(c ,fiber.StatusOK , 1 , "SUCCESS" , nil )
	return nil
}

func (s *HTTPServer) basicDetailsIncomeController(c *fiber.Ctx) error {
	//load basic details from request

	var basicDetails models.UserBasicDetailsIncome

	customErrors , err := ValidateRequest[models.UserBasicDetailsIncome](s, c , &basicDetails)
	if err!=nil{
		SendFullErrorResponse(c , http.StatusBadRequest , fmt.Errorf(constants.RequestError),customErrors)
		return nil
	}

	//basic details can be used here
	err = s.Srv.AddIncome(&basicDetails)
	if err != nil {
		errorResponse(c , http.StatusBadRequest , err)
		return nil
	}

	SendSuccessResponse(c ,fiber.StatusOK , 1 , "SUCCESS" , nil )
	return nil
}



func (s *HTTPServer) basicDetailsExpensesController(c *fiber.Ctx) error {
	//load basic details from request

	var basicDetails models.UserBasicDetailsExpenses

	customErrors , err := ValidateRequest[models.UserBasicDetailsExpenses](s, c , &basicDetails)
	if err!=nil{
		SendFullErrorResponse(c , http.StatusBadRequest , fmt.Errorf(constants.RequestError),customErrors)
		return nil
	}

	//basic details can be used here
	data , err := s.Srv.AddExpenses(&basicDetails)
	if err != nil {
		errorResponse(c , http.StatusBadRequest , err)
		return nil
	}



	SendSuccessResponse(c ,fiber.StatusOK , 1 , "SUCCESS" , data )
	return nil
}

func (s *HTTPServer) basicDetailsReportController(c *fiber.Ctx) error {
	//load basic details from request

	userId := c.Query("user_id")

	if userId == "" {
		errorResponse(c , http.StatusBadRequest , fmt.Errorf(constants.RequestError),)
	}

	data , err := s.Srv.GetUserInformation(userId)
	if err != nil {
		errorResponse(c , http.StatusBadRequest , err)
		return nil
	}

	SendSuccessResponse(c ,fiber.StatusOK , 1 , "SUCCESS" , data )
	return nil
}
