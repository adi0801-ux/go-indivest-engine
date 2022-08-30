package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"indivest-engine/services"
	"indivest-engine/utils"
)

// Server serves HTTP requests for this service

type HTTPServer struct {
	router       *fiber.App
	config       *utils.Config
	validator *validator.Validate
	Srv  *services.ServiceConfig
}

func (s *HTTPServer)RegisterRoutes(router *fiber.App) {

	router.Get("/", s.healthCheck)

	router.Post("/api/basicDetailsLanguage",s.basicDetailsLanguageController)

	router.Post("/api/basicDetailsIncome",s.basicDetailsIncomeController)

	router.Post("/api/basicDetailsExpenses",s.basicDetailsExpensesController)

	router.Get("/api/basicDetailsReport" , s.basicDetailsReportController)

}
func  (s *HTTPServer)HandleNotFound(router *fiber.App){

	// 404 Handler
	router.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}

// GetNewServer creates a new Http server and setup routing
func GetNewServer(
	Srv  *services.ServiceConfig,
	config *utils.Config) *HTTPServer {

	validate := validator.New()

	httpServer := &HTTPServer{config: config , validator: validate , Srv: Srv}

	router := fiber.New()

	// Add API Logger to Router
	//router.Use(utils.LoggerToFile(), recover.New())
	//router.Use(helmet.New())

	// Setup Routes here:
	httpServer.RegisterRoutes(router)


	httpServer.HandleNotFound(router)


	httpServer.router = router
	return httpServer
}

// StartServer Start the Gin Server at a specific address
func (s *HTTPServer) StartServer(a string) error {
	return s.router.Listen(a)
}


func (s *HTTPServer) healthCheck(c *fiber.Ctx) error {

	SendSuccessResponse(c ,fiber.StatusOK , 1 , "Alive!" , nil)
	return nil
}