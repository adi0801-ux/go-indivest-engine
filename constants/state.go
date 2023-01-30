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
	OnboardingsEndpoint             = "/onboardings"
	StartFullKycPrefix              = "/onboardings/"
	StartFullKycSuffix              = "/full_kyc"
	AddBankSuffix                   = "/bank_account"
	ReadPanCardSuffix               = "/read_pan_card"
	SubmitPanCardSuffix             = "/pan_card"
	UploadFile                      = "/upload_file"
	ReadAadhaarSuffix               = "/read_address_proof"
	SubmitAddressProof              = "address_proof"
	SubmitInvestorDetails           = "/form"
	UploadSignature                 = "/signature"
	UploadSelfie                    = "/selfie"
	StartVideoVerification          = "/start_video_verification"
	SubmitVideoVerification         = "/video_verification"
	GenerateKycContract             = "/generate_contract"
	ExecuteVerification             = "/execute_verification"
	India                           = "101"
	GetSip                          = "sips?"
	CreateSip                       = "sips"
	GetDeposits                     = "deposits?"
	CreateDeposit                   = "deposits"
	Payout                          = "N"
	Reinvest                        = "Y"
	Growth                          = "Z"
	Bonus                           = "B"
	CreateBasketOfDeposits          = "create_basket"
	RedirectURLAfterKyc             = "https://example.com/redirect"
	DefaultAMCCode                  = "MOF"
	ListAMCEndpoint                 = "/amcs"
	FundDetailsEndpoint             = "/funds"
	ShowAccounts                    = "/secure/accounts/"
	CreateWithdrawls                = "/secure/withdrawals/"
)

func GenerateFullKycURL(uuid string) string {
	return StartFullKycPrefix + uuid + StartFullKycSuffix
}

func GenerateAddBankURL(uuid string) string {

	return StartFullKycPrefix + uuid + AddBankSuffix
}

func GenerateAddPersonalDetailsURL(uuid string) string {

	return StartFullKycPrefix + uuid
}

func GenerateUploadFileURL(uuid string) string {

	return StartFullKycPrefix + uuid + UploadFile
}

func GenerateReadPanCardURL(uuid string) string {

	return StartFullKycPrefix + uuid + ReadPanCardSuffix
}

func GenerateSubmitPanCardURL(uuid string) string {

	return StartFullKycPrefix + uuid + SubmitPanCardSuffix
}

func GenerateReadAadharCardURL(uuid string) string {

	return StartFullKycPrefix + uuid + ReadAadhaarSuffix
}

func GenerateSubmitAadharCardURL(uuid string) string {

	return StartFullKycPrefix + uuid + SubmitAddressProof
}

func GenerateInvestorDetailsURL(uuid string) string {

	return StartFullKycPrefix + uuid + SubmitInvestorDetails
}

func GenerateUploadSignatureURL(uuid string) string {

	return StartFullKycPrefix + uuid + UploadSignature
}

func GenerateUploadSelfieURL(uuid string) string {

	return StartFullKycPrefix + uuid + UploadSelfie
}

func GenerateStartVideoVerificationURL(uuid string) string {
	return StartFullKycPrefix + uuid + StartVideoVerification
}

func GenerateSubmitVideoVerificationURL(uuid string) string {
	return StartFullKycPrefix + uuid + SubmitVideoVerification
}

func GenerateKYCContractURL(uuid string) string {
	return StartFullKycPrefix + uuid + GenerateKycContract
}

func GenerateKYCContractVerifyURL(uuid string) string {
	return StartFullKycPrefix + uuid + ExecuteVerification
}

func GenerateShowAccountsURL(uuid string) string {
	return ShowAccounts + uuid
}
