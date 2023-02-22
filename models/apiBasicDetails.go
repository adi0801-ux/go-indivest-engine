package models

type UserBasicDetailsLanguage struct {
	Language string `json:"language" validate:"required,oneof='english' 'hindi'"`
	UserId   string `json:"user_id"`
}

type UserBasicDetailsIncome struct {
	Income        float64 `json:"income" validate:"required,gt=0"`
	Age           string  `json:"age" validate:"required"`
	UserExpertise string  `json:"user_expertise"`
	Profession    string  `json:"profession" validate:"required,oneof='self_employed' 'salaried' 'professional' 'homemaker'"`
	UserId        string  `json:"user_id"`
}

type UserBasicDetailsExpenses struct {
	MonthlyEssentialExpense    float64 `json:"monthly_essential_expense" validate:"required,gte=0"`
	MonthlyNonEssentialExpense float64 `json:"monthly_non_essential_expense" validate:"required,gte=0"`
	MonthlyInvestments         float64 `json:"monthly_investments" validate:"gte=0"`
	UserId                     string  `json:"user_id"`
}

type CalculationResponse struct {
	InvestibleSurplus   InvestibleSurplus   `json:"investible_surplus"`
	CurrentPercentStats CurrentPercentStats `json:"current_percent_stats"`
	IdealPercentStats   IdealPercentStats   `json:"ideal_stats"`
	HealthSignal        string              `json:"health_signal"`
}
