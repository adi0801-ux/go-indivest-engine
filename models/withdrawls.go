package models

import "time"

type CreateWithdrawls struct {
	UserId               string    `json:"user_id"`
	Amount               float64   `json:"amount"`
	FundCode             time.Time `json:"fund_code"`
	PartnerTransactionId string    `json:"partner_transaction_id"`
}
type CreateWithdrawlAPI struct {
	Amount               float64   `json:"amount"`
	FundCode             time.Time `json:"fund_code"`
	AccountUuid          string    `json:"account_uuid"`
	PartnerTransactionId string    `json:"partner_transaction_id"`
}
type CreateWithdrawlAPIResponse struct {
}
