package models

import "time"

type UploadFileDB struct {
	ID     int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserId string `gorm:"column:user_id;not null" json:"user_id"`
	Url    string `gorm:"column:url" json:"url"`
}

type ReadPanCardDB struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time `gorm:"column:createdAt;not null;default:now()" json:"created_at"`
}

type StartFullKycDB struct {
	ID int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
}
