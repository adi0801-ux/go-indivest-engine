package models

import "time"

type CreateWithdrawals struct {
	UserId               string    `json:"user_id"`
	Amount               float64   `json:"amount"`
	FundCode             time.Time `json:"fund_code"`
	PartnerTransactionId string    `json:"partner_transaction_id"`
}
type CreateWithdrawalAPI struct {
	Widrawal struct {
		Amount               float64   `json:"amount"`
		FundCode             time.Time `json:"fund_code"`
		AccountUuid          string    `json:"account_uuid"`
		PartnerTransactionId string    `json:"partner_transaction_id"`
	} `json:"widrawal"`
}
type CreateWithdrawlAPIResponse struct {
}

// verifyWithdrawalOtp
type VerifyWithdrawalOtp struct {
	UserId string `json:"user_id"`
	Otp    string `json:"otp"`
}
type VerifyWithdrawalOtpAPI struct {
	Otp string `json:"otp"`
}
type VerifyWithdrawOtpAPIResponse struct {
}
