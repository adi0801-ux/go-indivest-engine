package models

//structure for buying of mutual fund

type BuyMutualFund struct {
	InvestmentType string  `json:"investment_type" validate:"required,oneof='SIP' 'One-Time'"`
	SchemeCode     string  `json:"scheme_code" validate:"required"`
	Amount         float64 `json:"amount" validate:"required,gt=0"`
	UserId         string  `json:"user_id"`
}

type RedeemMutualFund struct {
	SchemeCode string  `json:"scheme_code" validate:"required"`
	Amount     float64 `json:"amount" validate:"required,gt=0"`
	UserId     string  `json:"user_id"`
}

// Holdings structure for holding deatils
type Holdings struct {
	SchemeCode   string  `json:"scheme_code"`
	Units        float64 `json:"units"`
	CurrentValue float64 `json:"current_value"`
}

//structure for portfolio report
