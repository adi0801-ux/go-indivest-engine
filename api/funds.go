package api

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func (s *HTTPServer) fundHousesController(c *fiber.Ctx) error {
	//load basic details from request
	_, data, err := s.MfSrv.GetListOfFundHouses()
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return nil
	}

	SendSuccessResponse(c, fiber.StatusOK, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) fundDetailsController(c *fiber.Ctx) error {
	//load basic details from request
	_, data, err := s.MfSrv.GetListOfFunds()
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return nil
	}

	SendSuccessResponse(c, fiber.StatusOK, 1, "SUCCESS", data)
	return nil
}

func (s *HTTPServer) fundInfoController(c *fiber.Ctx) error {
	//load basic details from request
	AMFICode := c.Query("amfi_code")
	if AMFICode == "" {
		SendResponse(c, fiber.StatusBadRequest, 0, "enter amfi_code in params", nil, nil)
		return nil

	}
	_, data, err := s.MfSrv.GetFundDetail(AMFICode)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return nil
	}

	SendSuccessResponse(c, fiber.StatusOK, 1, "SUCCESS", data)
	return nil
}
