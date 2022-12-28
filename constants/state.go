package constants

const (
	RecommendedEmergencyFund        = 0.2
	RecommendedInvestibleFund       = 1 - RecommendedEmergencyFund
	HealthSignalRed                 = "RED"
	HealthSignalAmber               = "AMBER"
	HealthSignalGreen               = "GREEN"
	IdealMonthlyEssentialExpense    = 50
	IdealMonthlyNonEssentialExpense = 30
	IdealMonthlySavings             = 20
	OneTimePayment                  = "One-Time"
	SIP                             = "SIP"
	DefaultSandboxWalletHoldings    = 100000
	DefaultSIPActiveSatus           = 1
	DefaultfloatPrecissionAccepted  = 0.0005
	DefaultActivityDayLimit         = 6
	StartFullKyc                    = "full_kyc"
	UploadFile                      = "upload_file"
	ReadPanCard                     = "read_pan_card"
	SubmitPanCard                   = "pan_card"
	ReadAddressProof                = "read_address_proof"
	SubmitAddressProof              = "address_proof"
	SubmitInvestorDetails           = "form"
	UploadSignature                 = "signature"
	UploadSelfie                    = "selfie"
)
