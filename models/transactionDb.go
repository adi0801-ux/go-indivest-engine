package models

import (
	"time"
)

type CreateDepositsDb struct {
	FundCode                string    `gorm:"column:fund_code" json:"fund_code"`
	Amount                  float64   `gorm:"column:amount" json:"amount"`
	PaymentStatus           string    `gorm:"column:payment_status" json:"payment_status"`
	TransactionStatus       string    `gorm:"column:transaction_status" json:"transaction_status"`
	ID                      int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt               time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt               time.Time `gorm:"column:updated_at;not null;default:now()" json:"updated_at"`
	PaymentConfirmationTime string    `json:"payment_confirmation_time"`
	UserId                  string    `gorm:"column:user_id" json:"user_id"`
}

type CreateSipDb struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt     time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null;default:now()" json:"updated_at"`
	UserId        string    `gorm:"user_id" json:"user_id"`
	StartDate     string    `gorm:"start_date" json:"start_date"`
	EndDate       string    `gorm:"end_date" json:"end_date"`
	Frequency     string    `gorm:"frequency" json:"frequency"`
	FundCode      string    `gorm:"fund_code" json:"fund_code"`
	Amount        float64   `gorm:"amount" json:"amount"`
	Uuid          string    `gorm:"uuid" json:"uuid"`
	PaymentStatus string    `gorm:"payment_status" json:"payment_status"`
}
