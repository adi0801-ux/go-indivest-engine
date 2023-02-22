package models

import "time"

const TableNameSessionManager = "public.session_manager"

// SessionManager mapped from table <public.session_manager>
type SessionManager struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UserID      string    `gorm:"column:user_id;not null;unique" json:"user_id"`
	SessionID   string    `gorm:"column:session_id" json:"session_id"`
	SessionStep string    `gorm:"column:session_step" json:"session_step"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
}

// TableName SessionManager's table name
func (*SessionManager) TableName() string {
	return TableNameSessionManager
}

const TableNameUserDetails = "public.user_details"

// UserDetails mapped from table <public.user_details>
type UserDetails struct {
	ID                         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt                  time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UserID                     string    `gorm:"column:user_id;not null;unique" json:"user_id"`
	Language                   string    `gorm:"column:language" json:"language"`
	Age                        string    `gorm:"column:age" json:"age"`
	UserExpertise              string    `gorm:"column:user_expertise" json:"user_expertise"`
	GrossMonthlyIncome         float64   `gorm:"column:gross_monthly_income" json:"gross_monthly_income"`
	EducationalQualification   string    `gorm:"column:educational_qualification" json:"educational_qualification"`
	Profession                 string    `gorm:"column:profession" json:"profession"`
	Gender                     string    `gorm:"column:gender" json:"gender"`
	MonthlyEssentialExpense    float64   `gorm:"column:monthly_essential_expense" json:"monthly_essential_expense"`
	MonthlyNonEssentialExpense float64   `gorm:"column:monthly_non_essential_expense" json:"monthly_non_essential_expense"`
	MonthlySavings             float64   `gorm:"column:monthly_savings" json:"monthly_savings"`
	MonthlyInvestments         float64   `gorm:"column:monthly_investments" json:"monthly_investments"`
	MonthlyInvestibleSurplus   float64   `gorm:"column:monthly_investible_surplus" json:"monthly_investible_surplus"`
	UpdatedAt                  time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
}

// TableName UserDetails's table name
func (*UserDetails) TableName() string {
	return TableNameUserDetails
}

const TableNameUserReports = "public.user_reports"

// UserReports mapped from table <public.user_reports>
type UserReports struct {
	ID                      int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt               time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UserID                  string    `gorm:"column:user_id;not null;unique" json:"user_id"`
	EmergencyFund           float64   `gorm:"column:emergency_fund" json:"emergency_fund"`
	MonthlyInvestibleAmount float64   `gorm:"column:monthly_investible_amount" json:"monthly_investible_amount"`
	HealthSignal            string    `gorm:"column:health_signal" json:"health_signal"`
	UpdatedAt               time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
}

// TableName UserReports's table name
func (*UserReports) TableName() string {
	return TableNameUserReports
}
