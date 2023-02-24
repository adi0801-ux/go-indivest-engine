package models

import "time"

type FatcaCountryCode struct {
	Id          int       `gorm:"id, primaryKey"`
	Code        string    `gorm:"code"`
	Description string    `gorm:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()"`
}
type CountryCode struct {
	Id          int       `gorm:"id, primaryKey"`
	Code        string    `gorm:"code"`
	Description string    `gorm:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()"`
}
type GenderCodes struct {
	Id          int       `gorm:"id, primaryKey"`
	Code        string    `gorm:"code"`
	Description string    `gorm:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()"`
}
type Enums struct {
	Id          int       `gorm:"id, primaryKey"`
	Code        string    `gorm:"code"`
	Description string    `gorm:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()"`
}
type SourceOfWealth struct {
	Id          int       `gorm:"id, primaryKey"`
	Code        string    `gorm:"code"`
	Description string    `gorm:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()"`
}
type AddressType struct {
	Id          int       `gorm:"id, primaryKey"`
	Code        string    `gorm:"code"`
	Description string    `gorm:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()"`
}
type AnnualIncome struct {
	Id          int       `gorm:"id, primaryKey"`
	Code        string    `gorm:"code"`
	Description string    `gorm:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()"`
}
type ApplicationStatusCode struct {
	Id          int       `gorm:"id, primaryKey"`
	Code        string    `gorm:"code"`
	Description string    `gorm:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()"`
}
type MaritalStatusCode struct {
	Id          int       `gorm:"id, primaryKey"`
	Code        string    `gorm:"code"`
	Description string    `gorm:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()"`
}
type OccupationCode struct {
	Id          int       `gorm:"id, primaryKey"`
	Code        string    `gorm:"code"`
	Description string    `gorm:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()"`
}
