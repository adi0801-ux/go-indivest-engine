package models

import "mime/multipart"

// Fullkyc api models
type StartFullKyc struct {
	UserId             string `json:"user_id"`
	Name               string `json:"name"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phone_number"`
	FullKycRedirectUrl string `json:"full_kyc_redirect_url"`
}

type StartFullKycAPI struct {
	Name               string `json:"name"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phone_number"`
	FullKycRedirectUrl string `json:"full_kyc_redirect_url"`
}

// uploadFile api model
type UploadFile struct {
	UserId     string                `json:"user_id"`
	UploadFile *multipart.FileHeader `json:"upload_file"`
	Url        string                `json:"url"`
}

type UploadFileAPI struct {
	Url string `json:"string"`
}

// submitPanCard api model
type SubmitPanCard struct {
	Name        string `json:"name"`
	FathersName string `json:"fathers_name"`
	DateOfBirth string `json:"date_of_birth"`
	PanNumber   string `json:"pan_number"`
}
type SubmitPanCardAPI struct {
	Name        string `json:"name"`
	FathersName string `json:"fathers_name"`
	DateOfBirth string `json:"date_of_birth"`
	PanNumber   string `json:"pan_number"`
}
type SubmitPanCardAPIResponse struct {
}

// read pan card api models
type ReadPanCard struct {
	UserId   string `json:"user_id"`
	ImageUrl string `json:"image_url"`
}

type ReadPanCardAPI struct {
	ImageUrl string `json:"image_url"`
}

type ReadPanCardAPIResponse struct {
	OnBoarding  OnBoarding `json:"onboarding"`
	Name        string     `json:"name"`
	FathersName string     `json:"fathers_name"`
	DateOfBirth string     `json:"date_of_birth"`
	PanNumber   string     `json:"pan_number"`
}

type ReadAddressProof struct {
	UserId           string `json:"user_id"`
	AddressProofType string `json:"address_proof_type" validate:"required, oneof='aadhar','voter_id','passport','licence'"`
	ImageUrl         string `json:"image_url"`
}

type ReadAddressProofAPI struct {
	UserId           string `json:"user_id"`
	AddressProofType string `json:"address_proof_type" validate:"required, oneof='aadhar','voter_id','passport','licence'"`
	ImageUrl         string `json:"image_url"`
}

type ReadAddressProofAPIResponse struct {
	OnBoarding     OnBoarding `json:"on_boarding"`
	AadharUid      string     `json:"aadhar_uid"`
	LicenceNumber  string     `json:"licence_number"`
	PassportNumber string     `json:"passport_number"`
	VoterIdNumber  string     `json:"voter_id_number"`
	Name           string     `json:"name"`
	DateOfBirth    string     `json:"date_of_birth"`
	PinCode        string     `json:"pin_code"`
	Address        string     `json:"address"`
	District       string     `json:"district"`
	City           string     `json:"city"`
	State          string     `json:"state"`
	IssueDate      string     `json:"issue_date"`
	ExpiryDate     string     `json:"expiry_date"`
	FathersName    string     `json:"fathers_name"`
}

// submit addressProof models
type SubmitAddressProof struct {
	AddressProofType string `json:"address_proof_type"`
	Name             string `json:"name"`
	ExpiryDate       string `json:"expiry_date"`
	DateOfBirth      string `json:"date_of_birth"`
	IssueDate        string `json:"issue_date"`
	Address          string `json:"address"`
	City             string `json:"city"`
	State            string `json:"state"`
	District         string `json:"district"`
	PinCode          string `json:"pin_code"`
	LicenceNumber    string `json:"licence_number"`
	AadharUid        string `json:"aadhar_uid"`
	PassportNumber   string `json:"passport_number"`
	VoterIdNumber    string `json:"voter_id_number"`
}
type SubmitAddressProofAPI struct {
	AddressProofType string `json:"address_proof_type"`
	Name             string `json:"name"`
	ExpiryDate       string `json:"expiry_date"`
	DateOfBirth      string `json:"date_of_birth"`
	IssueDate        string `json:"issue_date"`
	Address          string `json:"address"`
	City             string `json:"city"`
	State            string `json:"state"`
	District         string `json:"district"`
	PinCode          string `json:"pin_code"`
	LicenceNumber    string `json:"licence_number"`
	AadharUid        string `json:"aadhar_uid"`
	PassportNumber   string `json:"passport_number"`
	VoterIdNumber    string `json:"voter_id_number"`
}
type SubmitAddressProofAPIResponse struct {
}

// submitInvestors Details
type SubmitInvestorDetails struct {
	Gender                       string `json:"gender"`
	MaritalStatus                string `json:"marital_status"`
	OccupationDescription        string `json:"occupation_description"`
	OccupationCode               string `json:"occupation_code"`
	CitizenshipCode              string `json:"citizenship_code"`
	CitizenshipCountry           string `json:"citizenship_country"`
	ApplicationStatusCode        string `json:"application_status_code"`
	ApplicationStatusDescription string `json:"application_status_description"`
	AnnualIncome                 string `json:"annual_income"`
}

type SubmitInvestorDetailsAPI struct {
	Gender                       string `json:"gender"`
	MaritalStatus                string `json:"marital_status"`
	OccupationDescription        string `json:"occupation_description"`
	OccupationCode               string `json:"occupation_code"`
	CitizenshipCode              string `json:"citizenship_code"`
	CitizenshipCountry           string `json:"citizenship_country"`
	ApplicationStatusCode        string `json:"application_status_code"`
	ApplicationStatusDescription string `json:"application_status_description"`
	AnnualIncome                 string `json:"annual_income"`
}
type SubmitInvestorDetailsAPIResponse struct {
}

// uploadSignature
type UploadSignature struct {
	UserId   string `json:"user_id"`
	ImageUrl string `json:"image_url"`
}

type UploadSignatureAPI struct {
	UserId   string `json:"user_id"`
	ImageUrl string `json:"image_url"`
}
type UploadSignatureAPIResponse struct {
}

// uploadSelfie added
type UploadSelfie struct {
	UserId   string `json:"user_id"`
	ImageUrl string `json:"image_url"`
}
type UploadSelfieAPI struct {
	UserId   string `json:"user_id"`
	ImageUrl string `json:"image_url"`
}
type UploadSelfieAPIResponse struct {
}

// startVideoVerification
type StartVideoVerification struct {
	UserId string `json:"user_id"`
}
type StartVideoVerificationAPI struct {
	UserId string `json:"user_id"`
}
type StartVideoVerificationAPIResponse struct {
	OnBoarding    OnBoarding `json:"on_boarding"`
	TransactionId string     `json:"transaction_id"`
	RandomNumber  string     `json:"random_number"`
}

// submitVideoVerification
type SubmitVideoVerification struct {
	UserId        string `json:"user_id"`
	VideoUrl      string `json:"video_url"`
	TransactionId string `json:"transaction_id"`
}
type SubmitVideoVerificationAPI struct {
	UserId        string `json:"user_id"`
	VideoUrl      string `json:"video_url"`
	TransactionId string `json:"transaction_id"`
}
type SubmitVideoVerificationAPIResponse struct {
}

// generateKycContract
type GenerateKycContract struct {
	UserId string `json:"user_id"`
}
type GenerateKycContractAPI struct {
	UserId string `json:"user_id"`
}
type GenerateKycContractAPIResponse struct {
	OnBoarding   OnBoarding `json:"on_boarding"`
	Url          string     `json:"url"`
	RandomNumber string     `json:"random_number"`
}

// onBoarding object struct
type OnBoarding struct {
	Uuid             string `json:"uuid"`
	PanNumber        string `json:"pan_number"`
	ExistingInvestor bool   `json:"existing_investor"`
	Name             string `json:"name"`
	DateOfBirth      string `json:"date_of_birth"`
	Email            string `json:"email"`
	PhoneNumber      string `json:"phone_number"`
	KycStatus        struct {
		Success string `json:"success"`
		Failure string `json:"failure"`
		Pending string `json:"pending"`
	} `json:"Kyc_status"`
	PanCardImageUrl      string `json:"pan_card_image_url"`
	FathersName          string `json:"fathers_name"`
	AddressProofImageUrl string `json:"address_proof_image_url"`
	AddressProofType     struct {
		Aadhar   string `json:"aadhar"`
		VoterId  string `json:"voter_id"`
		Passport string `json:"passport"`
		Licence  string `json:"licence"`
	} `json:"address_proof_type"`
	Address            string `json:"address"`
	City               string `json:"city"`
	Pincode            string `json:"pincode"`
	SignatureImageUrl  string `json:"signature_image_url"`
	SelfieImageUrl     string `json:"selfie_image_url"`
	CancelledChequeUrl string `json:"cancelled_cheque_url"`
	VideoUrl           string `json:"video_url"`
	AnnualIncome       string `json:"annual_income"`
	Gender             string `json:"gender"`
	Occupation         string `json:"occupation"`
	MaritalStatus      string `json:"marital_status"`
}
