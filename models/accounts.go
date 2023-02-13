package models

import "time"

type ShowAccount struct {
	UserId string `json:"user_id"`
}
type ShowAccountAPIResponse struct {
	AcntUuid string `json:"acnt_uuid"`
}

// ShowAccountDB gorm model
type ShowAccountDB struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UserId    string    `gorm:"column:user_id;not null" json:"user_id"`
	AmcId     string    `gorm:"column:amc_id" json:"amc_id"`
	AcntUuid  string    `gorm:"column:acnt_uuid" json:"acnt_uuid"`
}

type Webhook struct {
	Event       string      `json:"event"`
	SentAt      int         `json:"sent_at"`
	Payload     interface{} `json:"payload"`
	Signature   string      `json:"signature"`
	RedirectUrl string      `json:"redirect_url"`
}

type WebhooksAPI struct {
}

type WebhookDepositsCreate struct {
	Deposit struct {
		OnboardingUuid string `json:"onboarding_uuid"`
		AccountUuid    string `json:"account_uuid"`
		Uuid           string `json:"uuid"`
		Amount         int    `json:"amount"`
		Fund           struct {
			Name                       string `json:"name"`
			Active                     bool   `json:"active"`
			Code                       string `json:"code"`
			AmfiCode                   string `json:"amfi_code"`
			MinimumFirstTimeInvestment string `json:"minimum_first_time_investment"`
			MinimumOngoingInvestment   string `json:"minimum_ongoing_investment"`
			MinimumRedemptionAmount    string `json:"minimum_redemption_amount"`
			SettlementDays             int    `json:"settlement_days"`
			MinimumSipAmount           string `json:"minimum_sip_amount"`
			MinimumSwpAmount           string `json:"minimum_swp_amount"`
			MinimumStpAmount           string `json:"minimum_stp_amount"`
			FactsheetLink              string `json:"factsheet_link"`
			Category                   string `json:"category"`
			AmcId                      int    `json:"amc_id"`
			FundInfo                   struct {
				Nav         int `json:"nav"`
				ReturnYear1 int `json:"return_year_1"`
				ReturnYear3 int `json:"return_year_3"`
				ReturnYear5 int `json:"return_year_5"`
			} `json:"fund_info"`
			RiskRating   int    `json:"risk_rating"`
			ExpenseRatio string `json:"expense_ratio"`
			FundManagers string `json:"fund_managers"`
		} `json:"fund"`
		Status string `json:"status"`
	} `json:"deposit"`
}

type WebhookOnboardingCreate struct {
	PartnerTransactionId interface{} `json:"partner_transaction_id"`
	Uuid                 string      `json:"uuid"`
	AmcCode              string      `json:"amc_code"`
	ExistingInvestor     bool        `json:"existing_investor"`
	FullKycStatus        interface{} `json:"full_kyc_status"`
}

type WebhookAccountCreate struct {
	Account struct {
		OnboardingUuid string `json:"onboarding_uuid"`
		Uuid           string `json:"uuid"`
		AmcCode        string `json:"amc_code"`
	} `json:"account"`
}
