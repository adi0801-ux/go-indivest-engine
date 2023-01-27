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

// InvestmentAnalysis structure for portfolio report
type InvestmentAnalysis struct {
	SchemeCode          string  `json:"scheme_code"`
	Units               float64 `json:"units"`
	InvestedAmount      float64 `json:"invested_amount"`
	CurrentWorth        float64 `json:"current_worth"`
	PNL                 float64 `json:"pnl"`
	PNLPercentage       float64 `json:"pnl_percentage"`
	DayChange           float64 `json:"day_change"`
	DayChangePercentage float64 `json:"day_change_percentage"`
}

type UserMfActivity struct {
	CurrentWorth float64 `json:"current_worth"`
	Date         string  `json:"date"`
}

type UserMfInvestmentPanel struct {
	CurrentWorth    float64 `json:"current_worth"`
	TotalInvestment float64 `json:"total_investment"`
}

type InvestibleSurplus struct {
	EmergencyFund  float64 `json:"emergency_fund"`
	InvestibleFund float64 `json:"investible_fund"`
}

type CurrentPercentStats struct {
	EssentialExpenses    float64 `json:"essential_expenses"`
	NonEssentialExpenses float64 `json:"non_essential_expenses"`
	Savings              float64 `json:"savings"`
}

type IdealPercentStats struct {
	EssentialExpenses    float64 `json:"essential_expenses"`
	NonEssentialExpenses float64 `json:"non_essential_expenses"`
	Savings              float64 `json:"savings"`
}
