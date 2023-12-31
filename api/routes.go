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
	RiskSrv    *services.RiskCalculatorService
	SandboxSrv *services.SandboxServiceConfig
	MfSrv      *services.MFService
	UserSrv    *services.UserSrv
}

func (s *HTTPServer) RegisterRoutes(router *fiber.App) {
	rg := router.Group("/riskCalculator/api")
	rg.Get("/", s.healthCheck)
	rg.Use(s.AuthorizeMiddleware(s.config.AuthApi))
	{

		rg.Post("/basicDetailsLanguage", s.basicDetailsLanguageController)
		rg.Post("/basicDetailsIncome", s.basicDetailsIncomeController)
		rg.Post("/basicDetailsExpenses", s.basicDetailsExpensesController)
		rg.Get("/basicDetailsReport", s.basicDetailsReportController)
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

	mfEngine := router.Group("/mfEngine/api")

	mfEngine.Get("/occupationStatus", s.GetOccupationsController)
	mfEngine.Get("/genderCodes", s.GetGenderCodesController)
	mfEngine.Get("/martialStatus", s.GetMaritalStatusCodesController)
	mfEngine.Get("/countryCodes", s.GetCountryCodesController)
	mfEngine.Get("/annualIncome", s.GetAnnualIncomeLevelController)
	mfEngine.Get("/sourceOfWealth", s.GetSourceOfWealthController)
	mfEngine.Get("/fatcaCountryCode", s.GetFatcaCountryCodeController)
	mfEngine.Get("/applicationStatus", s.GetApplicationStatusController)
	mfEngine.Get("/addressType", s.GetAddressTypeController)

	mfEngine.Get("/", s.healthCheck)
	mfEngine.Use(s.AuthorizeMiddleware(s.config.AuthApi))
	{
		mfKyc := mfEngine.Group("/kyc")
		{

			mfKyc.Get("/status", s.CheckIfKycDoneController)
			mfKyc.Post("/start", s.StartFullKycController)
			mfKyc.Post("/addBank", s.AddBankAccountController)
			//to test

			// --
			mfKyc.Put("/addDetails", s.AddPersonalDetailsController)
			mfKyc.Post("/uploadPan", s.UploadPanCardController)
			mfKyc.Post("/verifyPan", s.SubmitPanCardController)
			//	to test
			mfKyc.Post("/uploadAadhaar", s.UploadAadhaarCardController)
			mfKyc.Post("/verifyAadhaar", s.SubmitAadhaarCardController)
			mfKyc.Post("/submitDetails", s.SubmitInvestorDetailsController)
			mfKyc.Post("/uploadSignature", s.UploadSignatureController)
			mfKyc.Post("/uploadSelfie", s.UploadSelfieController)
			mfKyc.Get("/startVideoVerification", s.StartVideoVerificationController)
			mfKyc.Post("/uploadVideoVerification", s.SubmitVideoVerificationController)
			mfKyc.Post("/signContract", s.GenerateKYCContractController)
			mfKyc.Post("/executeContract", s.ExecuteKYCVerificationController)
		}

		accounts := mfEngine.Group("/accounts")
		{
			accounts.Get("/show", s.ShowAccountDetailsController)
			accounts.Get("/holdings", s.GetHoldingsController)
			accounts.Get("/transactions", s.GetTransactionController)
			accounts.Get("/transactionsStatus", s.RequestStatusController)
			accounts.Get("/sortedTransaction", s.SortedTransactionController)

			watchList := accounts.Group("/watchList")
			{
				watchList.Post("/create", s.AddToWatchListController)
				watchList.Get("/show", s.ShowWatchListController)
			}

			withdrawals := accounts.Group("/withdrawals")
			{
				withdrawals.Post("/create", s.CreateWithdrawalController)
				withdrawals.Post("/verify_otp", s.VerifyWithdrawalOtpController)
			}
			deposits := accounts.Group("/deposits")
			{
				deposits.Get("/", s.GetDepositsController)
				deposits.Post("/create", s.CreateDepositsController)
			}
			sip := accounts.Group("/sips")
			{
				sip.Get("/", s.GetSipController)
				sip.Post("/create", s.CreateSipController)
			}
		}

		userInfo := mfEngine.Group("/user")
		userInfo.Use(s.AuthorizeMiddleware(s.config.AuthApi))
		{
			userInfo.Post("/onboardingQuestions", s.CreateOnBoardingQuestionsController)
		}
	}

	funds := router.Group("/funds/api")
	{
		funds.Get("/listFundHouses", s.fundHousesController)
		funds.Get("/fundDetails", s.fundDetailsController)
		funds.Get("/fundInfo", s.fundInfoController)
		funds.Post("/returnsCalculate", s.ReturnsInterestCalculatorController)
		funds.Get("/recommendation", s.RecommendationController)
		funds.Get("/popularFunds", s.PopularFundsController)
		funds.Get("/fundCategories", s.FundCategoriesController)
		funds.Get("/distinctCategories", s.DistinctCategoriesController)

	}

	webhook := router.Group("mfEngine/webhooks")
	webhook.Use(s.WebhookAuthenticationMiddleware())
	{
		webhook.Use("/savvy", s.ConnectWebhooksController)
	}

	authSrv := router.Group("/authService/api")
	{
		authSrv.Post("/user/signup", s.UserSignUpController)
		authSrv.Post("/user/login", s.UserLoginController)
		authSrv.Get("/token/verify", s.VerifyTokenController)
	}

}

//func (s *HTTPServer) HandleNotFound(router *fiber.App) {
//
//	// 404 Handler
//	router.Use(func(c *fiber.Ctx) error {
//		return c.SendStatus(404) // => 404 "Not Found"
//	})
//}

// GetNewServer creates a new Http server and setup routing
func GetNewServer(
	RiskSrv *services.RiskCalculatorService,
	SandBoxSrv *services.SandboxServiceConfig,
	MfSrv *services.MFService,
	UserSrv *services.UserSrv,

	config *utils.Config) *HTTPServer {

	validate := validator.New()

	httpServer := &HTTPServer{config: config, validator: validate, RiskSrv: RiskSrv, MfSrv: MfSrv, SandboxSrv: SandBoxSrv, UserSrv: UserSrv}

	router := fiber.New()

	// Add API Logger to Router
	router.Use(utils.LoggerToFile())
	//router.Use(recover.New())
	//router.Use(helmet.New())

	// Setup Routes here:
	httpServer.RegisterRoutes(router)

	//httpServer.HandleNotFound(router)

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
