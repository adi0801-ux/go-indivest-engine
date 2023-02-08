package models

// getDeposit api models
type GetDeposits struct {
	UserId string `json:"user_id"`
}
type GetDepositsAPI struct {
	AccountUuid string `json:"account_uuid"`
}

type GetDepositsAPIResponse struct {
}

// CreateDeposit api models
type CreateDeposit struct {
	Amount   string `json:"amount"`
	FundCode string `json:"fund_code"`
	UserId   string `json:"user_id"`
}

type CreateDepositAPI struct {
	Deposit struct {
		Amount             string `json:"amount"`
		FundCode           string `json:"fund_code"`
		AccountUuid        string `json:"account_uuid"`
		OnboardingUuid     string `json:"onboarding_uuid"`
		PaymentRedirectUrl string `json:"payment_redirect_url"`
	} `json:"deposit"`
}

type CreateDepositAPIResponse struct {
	Deposit Deposit `json:"deposit"`
	Url     string  `json:"url"`
}

type CreateDepositLocal struct {
	Deposit Deposit `json:"deposit"`
	Url     string  `json:"url"`
}

// createBasketOfDeposits
type CreateBasketOfDeposits struct {
	PaymentRedirectUrl   string    `json:"payment_redirect_url"`
	AccountUuid          string    `json:"account_uuid"`
	OnBoardingUuid       string    `json:"onBoarding_uuid"`
	PartnerTransactionId string    `json:"partner_transaction_id"`
	DepositsParts        [2]string `json:"deposits_parts"`
}
type CreateBasketOfDepositsAPI struct {
	PaymentRedirectUrl   string    `json:"payment_redirect_url"`
	AccountUuid          string    `json:"account_uuid"`
	OnBoardingUuid       string    `json:"onBoarding_uuid"`
	PartnerTransactionId string    `json:"partner_transaction_id"`
	DepositsParts        [2]string `json:"deposits_parts"`
}
type CreateBasketOfDepositsAPIResponse struct {
	Deposit Deposit `json:"deposit"`
	Url     string  `json:"url"`
}

// OBJECT Deposit
type Deposit struct {
	Uuid                   string  `json:"uuid"`
	FundCode               string  `json:"fund_code"`
	FundName               string  `json:"fund_name"`
	Amount                 float64 `json:"amount"`
	CurrentAmount          float64 `json:"current_amount"`
	Units                  float64 `json:"units"`
	Status                 string  `json:"status" validate:"required oneof:'created','payment_mode', 'submitted_to_rta', 'completed','error'"`
	StatusDescription      string  `json:"status_description"`
	ReinvestMode           string  `json:"reinvest_mode" validate:"required oneof: 'Payout','Reinvest','Growth','Bonus'"`
	PartnerTransactionId   string  `json:"partner_transaction_id"`
	UserCompletedPaymentAt string  `json:"user_completed_payment_at"`
	TransferredToAmcAt     string  `json:"transferred_to_amc_at"`
	CreatedAt              string  `json:"created_at"`
	SipUuid                string  `json:"sip_uuid"`
	StpUuid                string  `json:"stp_uuid"`
}

// Withdrawals
type CreateWithdrawals struct {
	UserId   string `json:"user_id"`
	Amount   string `json:"amount"`
	FundCode string `json:"fund_code"`
}

type CreateWithdrawalAPI struct {
	Withdrawal struct {
		Amount               string `json:"amount"`
		FundCode             string `json:"fund_code"`
		AccountUuid          string `json:"account_uuid"`
		PartnerTransactionId string `json:"partner_transaction_id"`
	} `json:"withdrawal"`
}

type CreateWithdrawlAPIResponse struct {
}

type CreateWithdrawalLocal struct {
}

// verifyWithdrawalOtp
type VerifyWithdrawalOtp struct {
	UserId string `json:"user_id"`
	Otp    string `json:"otp"`
}
type VerifyWithdrawalOtpAPI struct {
	Withdrawal struct {
		Otp string `json:"otp"`
	} `json:"withdrawal"`
}
type VerifyWithdrawOtpAPIResponse struct {
}

//SIP models

// GetSip models
type GetSip struct {
	UserId string `json:"user_id"`
}

type GetSipAPI struct {
	AccountUuid string `json:"account_uuid"`
}
type GetSipAPIResponse struct {
}

// CreateSip models

type CreateSip struct {
	Amount    float64 `json:"amount"`
	FundCode  string  `json:"fund_code"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
	UserId    string  `json:"user_id"`
}

type CreateSipAPI struct {
	Sip struct {
		Amount               float64 `json:"amount"`
		FundCode             string  `json:"fund_code"`
		AccountUuid          string  `json:"account_uuid"`
		OnboardingUuid       string  `json:"onboarding_uuid"`
		PartnerTransactionId string  `json:"partner_transaction_id"`
		StartDate            string  `json:"start_date"`
		EndDate              string  `json:"end_date"`
		Frequency            string  `json:"frequency"`
		MandateRedirectUrl   string  `json:"mandate_redirect_url"`
	} `json:"sip"`
}

type CreateSipApiResponse struct {
	Success bool   `json:"success"`
	Url     string `json:"url"`
	Sip     struct {
		Uuid          string        `json:"uuid"`
		Amount        float64       `json:"amount"`
		FundCode      string        `json:"fund_code"`
		FundName      string        `json:"fund_name"`
		MandateStatus interface{}   `json:"mandate_status"`
		PaymentLink   interface{}   `json:"payment_link"`
		StartDate     string        `json:"start_date"`
		EndDate       string        `json:"end_date"`
		Frequency     string        `json:"frequency"`
		Active        interface{}   `json:"active"`
		CreatedAt     string        `json:"created_at"`
		Deposits      []interface{} `json:"deposits"`
		Fund          struct {
			Name                       string      `json:"name"`
			Active                     bool        `json:"active"`
			Code                       string      `json:"code"`
			AmfiCode                   interface{} `json:"amfi_code"`
			MinimumFirstTimeInvestment string      `json:"minimum_first_time_investment"`
			MinimumOngoingInvestment   string      `json:"minimum_ongoing_investment"`
			MinimumRedemptionAmount    string      `json:"minimum_redemption_amount"`
			SettlementDays             int         `json:"settlement_days"`
			MinimumSipAmount           string      `json:"minimum_sip_amount"`
			MinimumSwpAmount           string      `json:"minimum_swp_amount"`
			MinimumStpAmount           string      `json:"minimum_stp_amount"`
			FactsheetLink              string      `json:"factsheet_link"`
			Category                   string      `json:"category"`
			AmcId                      int         `json:"amc_id"`
			FundInfo                   struct {
				Nav         int     `json:"nav"`
				ReturnYear1 float64 `json:"return_year_1"`
				ReturnYear3 float64 `json:"return_year_3"`
				ReturnYear5 float64 `json:"return_year_5"`
			} `json:"fund_info"`
			RiskRating   interface{} `json:"risk_rating"`
			ExpenseRatio interface{} `json:"expense_ratio"`
			FundManagers interface{} `json:"fund_managers"`
		} `json:"fund"`
		MandateGateway string `json:"mandate_gateway"`
	} `json:"sip"`
}

type CreateSipLocal struct {
}

type Holding struct {
	FundCode string `json:"fund_code"`
	AmcCode  string `json:"amc_code"`
}
type HoldingApi struct {
	FundCode string `json:"fund_code"`
	AmcCode  string `json:"amc_code"`
}