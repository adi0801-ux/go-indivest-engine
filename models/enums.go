package models

type FatcaCountryCode struct {
	Id          int    `gorm:"id"`
	Code        string `gorm:"code"`
	Description string `gorm:"description"`
}
type CountryCode struct {
	Id          int    `gorm:"id"`
	Code        string `gorm:"code"`
	Description string `gorm:"description"`
}

type GenderCodes struct {
	Id          int    `gorm:"id"`
	Code        string `gorm:"code"`
	Description string `gorm:"description"`
}
type Enums struct {
	Id          int    `gorm:"id"`
	Code        string `gorm:"code"`
	Description string `gorm:"description"`
}
type SourceOfWealth struct {
	Id          int    `gorm:"id"`
	Code        string `gorm:"code"`
	Description string `gorm:"description"`
}
type AddressType struct {
	Id          int    `gorm:"id"`
	Code        string `gorm:"code"`
	Description string `gorm:"description"`
}
type AnnualIncome struct {
	Id          int    `gorm:"id"`
	Code        string `gorm:"code"`
	Description string `gorm:"description"`
}
type ApplicationStatusCode struct {
	Id          int    `gorm:"id"`
	Code        string `gorm:"code"`
	Description string `gorm:"description"`
}
type MaritalStatusCode struct {
	Id          int    `gorm:"id"`
	Code        string `gorm:"code"`
	Description string `gorm:"description"`
}
type OccupationCode struct {
	Id          int    `gorm:"id"`
	Code        string `gorm:"code"`
	Description string `gorm:"description"`
}
