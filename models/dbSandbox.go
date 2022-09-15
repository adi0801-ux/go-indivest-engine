package models

import "time"

const TableNameUserWallet = "public.user_mutual_fund_wallet"

// UserWallet mapped from table <public.user_wallet>
type UserWallet struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UserID    string    `gorm:"column:user_id;not null;unique" json:"user_id"`
	INR       float64   `gorm:"column:inr" json:"inr"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
}

// TableName UserWallet's table name
func (*UserWallet) TableName() string {
	return TableNameUserWallet
}

const TableNameUserMFHoldings = "public.user_mutual_fund_holdings"

// UserMFHoldings mapped from table <public.user_mutual_fund_holdings>
type UserMFHoldings struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UserID     string    `gorm:"column:user_id;not null" json:"user_id"`
	SchemeCode string    `gorm:"column:scheme_code;not null" json:"scheme_code"`
	FundUnits  float64   `gorm:"column:units" json:"units"`
	UpdatedAt  time.Time `gorm:"column:updated_at;default:now()" json:"updated_at"`
}

// TableName UserMFHoldings's table name
func (*UserMFHoldings) TableName() string {
	return TableNameUserMFHoldings
}

const TableNameUserMFTransactions = "public.user_mutual_fund_transactions"

// UserMFTransactions mapped from table <public.user_mutual_fund_transactions>
type UserMFTransactions struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TransactionID   string    `gorm:"column:transaction_id;not null;unique" json:"transaction_id"`
	SipID           string    `gorm:"column:sip_id;not null" json:"sip_id"`
	CreatedAt       time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UserID          string    `gorm:"column:user_id;not null" json:"user_id"`
	SchemeCode      string    `gorm:"column:scheme_code;not null" json:"scheme_code"`
	FundUnits       float64   `gorm:"column:units;not null" json:"units"`
	NAV             float64   `gorm:"column:nav;not null" json:"nav"`
	INRAmount       float64   `gorm:"column:inr_amount;not null" json:"inr_amount"`
	TransactionType string    `gorm:"column:transaction_type;not null" json:"transaction_type"` // buy /sell
	InvestmentType  string    `gorm:"column:investment_type;not null" json:"investment_type"`
}

// TableName UserMFTransactions's table name
func (*UserMFTransactions) TableName() string {
	return TableNameUserMFTransactions
}

const TableNameUserMFActiveSIP = "public.user_mutual_fund_active_sip"

// UserMFActiveSIP mapped from table <public.user_mutual_fund_active_sip>
type UserMFActiveSIP struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UserID     string    `gorm:"column:user_id;not null" json:"user_id"`
	SipID      string    `gorm:"column:sip_id;not null" json:"sip_id"`
	SchemeCode string    `gorm:"column:scheme_code;not null" json:"scheme_code"`
	SIPAmount  float64   `gorm:"column:sip_amount;not null" json:"sip_amount"`
	SIPDate    int       `gorm:"column:sip_date;not null" json:"sip_date"`
	Active     int       `gorm:"column:active;not null" json:"active"`
}

// TableName UserMFVirtualPortfolio's table name
func (*UserMFActiveSIP) TableName() string {
	return TableNameUserMFActiveSIP
}

const TableNameUserMFDailyReport = "public.user_mutual_fund_daily_report"

// UserMFDailyReport mapped from table <public.user_mutual_fund_daily_report>
type UserMFDailyReport struct {
	ID              int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt       time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UserID          string    `gorm:"column:user_id;not null" json:"user_id"`
	SchemeCode      string    `gorm:"column:scheme_code;not null" json:"scheme_code"`
	InvestmentWorth float64   `gorm:"column:investment_worth;not null" json:"investment_worth"`
}

// TableName UserMFDailyReport's table name
func (*UserMFDailyReport) TableName() string {
	return TableNameUserMFDailyReport
}
