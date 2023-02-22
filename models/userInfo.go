package models

import "time"

type UserSignup struct {
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	EmailId     string `json:"email_id" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
	DeviceToken string `json:"device_token"`
	DeviceType  string `json:"device_type"`
}

type UserLogin struct {
	EmailId  string `json:"email_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt     time.Time `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null;default:now()" json:"updated_at"`
	UserId        string    `gorm:"user_id" json:"user_id"`
	FirstName     string    `gorm:"first_name" json:"first_name"`
	LastName      string    `gorm:"last_name" json:"last_name"`
	EmailId       string    `gorm:"email_id" json:"email_id"`
	PhoneNumber   string    `gorm:"phone_number" json:"phone_number"`
	Password      string    `gorm:"password" json:"password"`
	DeviceToken   string    `gorm:"device_token" json:"device_token"`
	DeviceType    string    `gorm:"device_type" json:"device_type"`
	ProfileStatus string    `gorm:"profile_status" json:"profile_status"`
}

func (*User) TableName() string {
	return "user_info"

}

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
