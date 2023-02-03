package models

import "time"

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
	Amount               string    `json:"amount"`
	FundCode             time.Time `json:"fund_code"`
	PaymentRedirectUrl   string    `json:"payment_redirect_url"`
	AccountUuid          string    `json:"account_uuid"`
	OnBoardingUuid       string    `json:"on_boarding_uuid"`
	PartnerTransactionId string    `json:"partner_transaction_id"`
	UserId               string    `json:"user_id"`
}
type CreateDepositAPI struct {
	Amount               string    `json:"amount"`
	FundCode             time.Time `json:"fund_code"`
	PaymentRedirectUrl   string    `json:"payment_redirect_url"`
	AccountUuid          string    `json:"account_uuid"`
	OnBoardingUuid       string    `json:"on_boarding_uuid"`
	PartnerTransactionId string    `json:"partner_transaction_id"`
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
	Uuid                   string    `json:"uuid"`
	FundCode               string    `json:"fund_code"`
	FundName               string    `json:"fund_name"`
	Amount                 int64     `json:"amount"`
	CurrentAmount          float64   `json:"current_amount"`
	Units                  float64   `json:"units"`
	Status                 string    `json:"status" validate:"required oneof:'created','payment_mode', 'submitted_to_rta', 'completed','error'"`
	StatusDescription      string    `json:"status_description"`
	ReinvestMode           string    `json:"reinvest_mode" validate:"required oneof: 'Payout','Reinvest','Growth','Bonus'"`
	PartnerTransactionId   string    `json:"partner_transaction_id"`
	UserCompletedPaymentAt time.Time `json:"user_completed_payment_at"`
	TransferredToAmcAt     time.Time `json:"transferred_to_amc_at"`
	CreatedAt              time.Time `json:"created_at"`
	SipUuid                string    `json:"sip_uuid"`
	StpUuid                string    `json:"stp_uuid"`
}
