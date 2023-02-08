package models

import "time"

type ShowAccount struct {
	UserId string `json:"user_id"`
}
type ShowAccountAPIResponse struct {
	AcntUuid string `json:"acnt_uuid"`
}

// ShowAccountDB gorm model
type ShowAccountDB struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UserId    string    `gorm:"column:user_id;not null" json:"user_id"`
	Uuid      string    `gorm:"column:uuid" json:"uuid"`
	AcntUuid  string    `gorm:"column:acnt_uuid" json:"acnt_uuid"`
}

type Webhook struct {
	Event       string      `json:"event"`
	SentAt      string      `json:"sent_at"`
	Payload     interface{} `json:"payload"`
	Signature   string      `json:"signature"`
	RedirectUrl string      `json:"redirect_url"`
}

type WebhooksAPI struct {
}
