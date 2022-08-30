package models

type InvestibleSurplus struct {
	EmergencyFund float64 `json:"emergency_fund"`
	InvestibleFund float64 `json:"investible_fund"`
}

type CurrentPercentStats struct {
	EssentialExpenses float64 `json:"essential_expenses"`
	NonEssentialExpenses float64 `json:"non_essential_expenses"`
	Savings float64 `json:"savings"`
}

type IdealPercentStats struct {
	EssentialExpenses float64 `json:"essential_expenses"`
	NonEssentialExpenses float64 `json:"non_essential_expenses"`
	Savings float64 `json:"savings"`
}