package models

import "time"

type UserQuestioner struct {
	InvestingInterest string `json:"investing_interest"`
	Profession        string `json:"profession"`
	UserId            string `json:"user_id"`
}

type UserLeads struct {
	ID                int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt         time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at;not null;default:now()" json:"updated_at"`
	UserId            string    `gorm:"user_id" json:"user_id"`
	InvestingInterest string    `gorm:"investing_interest" json:"investing_interest"`
	Profession        string    `gorm:"profession"  json:"profession"`
}
