package models

import "time"

type UploadFileDB struct {
	ID     int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserId string `gorm:"column:user_id;not null" json:"user_id"`
	Url    string `gorm:"column:url" json:"url"`
}

type ReadPanCardDB struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time `gorm:"column:createdAt;not null;default:now()" json:"created_at"`
	UserId      string    `gorm:"column:user_id;not null" json:"user_id"`
	Name        string    `gorm:"column:name" json:"name"`
	FathersName string    `gorm:"column:fathers_name" json:"fathers_name"`
	DateOfBirth string    `gorm:"column:date_of_birth" json:"date_of_birth"`
	PanNumber   string    `gorm:"column:pan_number" json:"pan_number"`
}

type StartFullKycDB struct {
	ID int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
}

type ReadAddressProofDB struct {
	ID             int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt      time.Time  `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	OnBoarding     OnBoarding `gorm:"column:on_boarding" json:"on_boarding"`
	AadharUid      string     `gorm:"column:aadhar_uid" json:"aadhar_uid"`
	AadharVid      string     `gorm:"column:aadhar_vid" json:"aadhar_vid"`
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
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	TransactionId string    `gorm:"column:transaction_id" json:"transaction_id"`
	RandomNumber  string    `gorm:"column:random_number" json:"random_number"`
	UserId        string    `gorm:"column:user_id" json:"user_id"`
}

type GenerateKycContractDB struct {
	ID           int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt    time.Time  `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	OnBoarding   OnBoarding `gorm:"column:on_boarding" json:"on_boarding"`
	Url          string     `gorm:"column:url" json:"url"`
	RandomNumber string     `gorm:"column:random_number" json:"random_number"`
	UserId       string     `gorm:"column:user_id" json:"user_id"`
}

type OnboardingObjectDB struct {
	ID                   int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt            time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UserId               string    `gorm:"column:user_id;not null" json:"user_id"`
	Uuid                 string    `gorm:"column:uuid" json:"uuid"`
	PanNumber            string    `gorm:"column:pan_number" json:"pan_number"`
	ExistingInvestor     string    `gorm:"column:existing_investor" json:"existing_investor"`
	Name                 string    `gorm:"column:name" json:"name"`
	Email                string    `gorm:"column:email" json:"email"`
	PhoneNumber          string    `gorm:"column:phone_number" json:"phone_number"`
	DateOfBirth          string    `gorm:"column:date_of_birth" json:"date_of_birth"`
	KycStatus            string    `gorm:"column:kyc_status" json:"kyc_status"`
	PanCardImageUrl      string    `gorm:"column:pan_card_image_url" json:"pan_card_image_url"`
	FathersName          string    `gorm:"column:fathers_name" json:"fathers_name"`
	AddressProofImageUrl string    `gorm:"column:address_proof_image_url" json:"address_proof_image_url"`
	AddressProofType     string    `gorm:"column:address_proof_type" json:"address_proof_type"`
	Address              string    `gorm:"column:address" json:"address"`
	City                 string    `gorm:"column:city" json:"city"`
	Pincode              string    `gorm:"column:pincode" json:"pincode"`
	SignatureImageUrl    string    `gorm:"column:signature_image_url" json:"signature_image_url"`
	SelfieImageUrl       string    `gorm:"column:selfie_image_url" json:"selfie_image_url"`
	CancelledChequeUrl   string    `gorm:"column:cancelled_cheque_url" json:"cancelled_cheque_url"`
	VideoUrl             string    `gorm:"column:video_url" json:"video_url"`
	AnnualIncome         string    `gorm:"column:annual_income" json:"annual_income"`
	Gender               string    `gorm:"column:gender" json:"gender"`
	Occupation           string    `gorm:"column:occupation" json:"occupation"`
	MaritalStatus        string    `gorm:"column:marital_status" json:"marital_status"`
}

func (o *OnboardingObjectDB) TableName() string {
	return "onboarding_object"

}

type BankAccountDB struct {
	ID            int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	UserId        string    `gorm:"column:user_id;not null" json:"user_id"`
	AccountNumber string    `gorm:"column:account_number" json:"account_number"`
	BankName      string    `gorm:"column:bank_name" json:"bank_name"`
	BranchName    string    `gorm:"column:branch_name" json:"branch_name"`
	BankCity      string    `gorm:"column:bank_city" json:"bank_city"`
	IfscCode      string    `gorm:"column:ifsc_code" json:"ifsc_code"`
}

type OccupationDB struct {
	ID             int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	OccupationId   string    `gorm:"column:occupation_id;not null" json:"occupation_id"`
	OccupationName string    `gorm:"column:occupation_name" json:"occupation_name"`
}

type GenderCodesDB struct {
	ID         int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	GenderCode string    `gorm:"column:gender_code;not null" json:"gender_code"`
	GenderName string    `gorm:"column:gender_name" json:"gender_name"`
}

type MaritalStatusCodesDB struct {
	ID                int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt         time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	MaritalStatus     string    `gorm:"column:martial_status;not null" json:"martial_status"`
	MaritalStatusName string    `gorm:"column:martial_status_name" json:"martial_status_name"`
}

type CountryCodesDB struct {
	ID          int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	CountryCode string    `gorm:"column:country_code;not null" json:"country_code"`
	CountryName string    `gorm:"column:country_name" json:"country_name"`
}

type IncomeLevelDB struct {
	ID               int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;default:now()" json:"created_at"`
	AnualIncomeLevel string    `gorm:"column:income_level;not null" json:"income_level"`
}
