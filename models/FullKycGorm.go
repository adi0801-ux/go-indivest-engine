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

type ReadAddressProofDB struct {
	ID             int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt      time.Time  `gorm:"column:createdAt;not null;default:now()" json:"created_at"`
	OnBoarding     OnBoarding `gorm:"column:onboarding" json:"on_boarding"`
	AadharUid      string     `gorm:"column:aadhar_uid" json:"aadhar_uid"`
	LicenceNumber  string     `gorm:"column:licence_number" json:"licence_number"`
	PassportNumber string     `gorm:"column:passport_number" json:"passport_number"`
	VoterIdNumber  string     `gorm:"column:voter_id_number" json:"voter_id_number"`
	Name           string     `gorm:"column:name" json:"name"`
	DateOfBirth    string     `gorm:"column:date_of_birth" json:"date_of_birth"`
	PinCode        string     `gorm:"column:pin_code" json:"pin_code"`
	Address        string     `gorm:"column:address" json:"address"`
	District       string     `gorm:"column:district" json:"district"`
	City           string     `gorm:"column:city" json:"city"`
	State          string     `gorm:"column:state" json:"state"`
	IssueDate      string     `gorm:"issue_date" json:"issue_date"`
	ExpiryDate     string     `gorm:"column:expiry_date" json:"expiry_date"`
	FathersName    string     `gorm:"fathers_name" json:"fathers_name"`
}

type StartVideoVerificationDB struct {
	ID            int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt     time.Time  `gorm:"column:createdAt;not null;default:now()" json:"created_at"`
	OnBoarding    OnBoarding `json:"on_boarding"`
	TransactionId string     `json:"transaction_id"`
	RandomNumber  string     `json:"random_number"`
}
