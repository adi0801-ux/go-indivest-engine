package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"indivest-engine/services"
	"indivest-engine/utils"
)

// Server serves HTTP requests for this service

type HTTPServer struct {
	router     *fiber.App
	config     *utils.Config
	validator  *validator.Validate
	Srv        *services.ServiceConfig
	SandboxSrv *services.SandboxServiceConfig
}

func (s *HTTPServer) RegisterRoutes(router *fiber.App) {
	rg := router.Group("/mfEngine")
	rg.Get("/", s.healthCheck)
	rg.Use(s.AuthorizeMiddleware(s.config.AuthApi))
	{

		rg.Post("/api/basicDetailsLanguage", s.basicDetailsLanguageController)

		rg.Post("/api/basicDetailsIncome", s.basicDetailsIncomeController)

		rg.Post("/api/basicDetailsExpenses", s.basicDetailsExpensesController)

		rg.Get("/api/basicDetailsReport", s.basicDetailsReportController)
	}

	sandbox := router.Group("/mfSandbox/api")
	sandbox.Get("/", s.healthCheck)
	sandbox.Use(s.AuthorizeMiddleware(s.config.AuthApi))
	{
		//sandbox.Get("/", s.healthCheck)

		sandbox.Post("/buyMutualFund", s.sandboxBuyMutualFund)

		sandbox.Get("/holding", s.sandboxGetHolding)

		sandbox.Get("/allHoldings", s.sandboxGetAllHolding)

		sandbox.Get("/wallet", s.sandboxGetWallet)

		sandbox.Get("/transactions", s.sandboxGetTransactions)

		sandbox.Post("/redeemMutualFund", s.sandboxRedeemMutualFund)

		sandbox.Get("/investmentAnalysis", s.sandboxUserInvestmentAnalysis)

		sandbox.Get("/userMfActivity", s.sandboxUserMfActivity)

		sandbox.Get("/userInvestmentPanel", s.sandboxUserMfInvestmentPanel)

	}
}
func (s *HTTPServer) HandleNotFound(router *fiber.App) {

	// 404 Handler
	router.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}

// GetNewServer creates a new Http server and setup routing
func GetNewServer(
	Srv *services.ServiceConfig,
	SandBoxSrv *services.SandboxServiceConfig,

	config *utils.Config) *HTTPServer {

	validate := validator.New()

	httpServer := &HTTPServer{config: config, validator: validate, Srv: Srv, SandboxSrv: SandBoxSrv}

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
	SendSuccessResponse(c, fiber.StatusOK, 1, "Alive!", nil)
	return nil
}
