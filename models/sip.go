package models

import "time"

// GetSip models
type GetSip struct {
	AccountUuid string `json:"account_uuid"`
}

type GetSipAPI struct {
	AccountUuid string `json:"account_uuid"`
}
type GetSipAPIResponse struct {
}

// CreateSip models
type CreateSip struct {
	Amount               float64   `json:"amount"`
	FundCode             time.Time `json:"fund_code"`
	AccountUuid          string    `json:"account_uuid"`
	OnBoardingUuid       string    `json:"on_boarding_uuid"`
	PartnerTransactionId string    `json:"partner_transaction_id"`
	StartDate            time.Time `json:"start_date"`
	EndDate              time.Time `json:"end_date"`
	Frequency            string    `json:"frequency" validate:"required, oneof='monthly','weekly','daily','ad-hoc'"`
	MandateRedirectUrl   string    `json:"mandate_redirect_url"`
}
type CreateSipAPI struct {
	Amount               float64   `json:"amount"`
	FundCode             time.Time `json:"fund_code"`
	AccountUuid          string    `json:"account_uuid"`
	OnBoardingUuid       string    `json:"on_boarding_uuid"`
	PartnerTransactionId string    `json:"partner_transaction_id"`
	StartDate            time.Time `json:"start_date"`
	EndDate              time.Time `json:"end_date"`
	Frequency            string    `json:"frequency" validate:"required, oneof='monthly','weekly','daily','ad-hoc'"`
	MandateRedirectUrl   string    `json:"mandate_redirect_url"`
}
type CreateSipAPIResponse struct {
}

type CreateSipLocal struct {
}
