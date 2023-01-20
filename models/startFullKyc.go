package models

import "mime/multipart"

type CheckKYCUser struct {
	PanNumber string `json:"pan_number"`
	UserId    string `json:"user_id"`
}

type CheckKYCUserAPI struct {
	Onboarding struct {
		PanNumber string `json:"pan_number"`
		AmcCode   string `json:"amc_code"`
	} `json:"onboarding"`
}

type AddBankAccount struct {
	AccountNumber string `json:"account_number"`
	UserId        string `json:"user_id"`
	IFSC          string `json:"ifsc"`
}

type AddBankAccountAPI struct {
	Onboarding struct {
		AccountNumber string `json:"account_number"`
		IfscCode      string `json:"ifsc_code"`
	} `json:"onboarding"`
}

type AddPersonalDetails struct {
	Address     string `json:"address"`
	City        string `json:"city"`
	Pincode     string `json:"pincode"`
	DateOfBirth string `json:"date_of_birth"`
	Occupation  string `json:"occupation"`
	UserId      string `json:"user_id"`
}

type AddPersonalDetailsAPI struct {
	Onboarding struct {
		Address     string `json:"address"`
		City        string `json:"city"`
		Pincode     string `json:"pincode"`
		DateOfBirth string `json:"date_of_birth"`
		Occupation  string `json:"occupation"`
	} `json:"onboarding"`
}
type AddBankAccountAPIResponse struct {
	Onboarding  OnBoarding  `json:"onboarding"`
	BankAccount BankAccount `json:"bank_account"`
}

type BankAccount struct {
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	BranchName    string `json:"branch_name"`
	BankCity      string `json:"bank_city"`
	IfscCode      string `json:"ifsc_code"`
}

type UploadPanCard struct {
	PanCard *multipart.FileHeader `json:"pan_card"`
	UserId  string                `json:"user_id"`
}

// Fullkyc api models
type StartFullKyc struct {
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type StartFullKycAPI struct {
	Onboarding struct {
		Email              string `json:"email"`
		Name               string `json:"name"`
		PhoneNumber        string `json:"phone_number"`
		FullKycRedirectUrl string `json:"full_kyc_redirect_url"`
	} `json:"onboarding"`
}

// uploadFile api model
type UploadFile struct {
	UUID       string                `json:"uuid"`
	UploadFile *multipart.FileHeader `json:"upload_file"`
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
	UserId      string `json:"user_id"`
}
type SubmitPanCardAPI struct {
	Onboarding struct {
		Name        string `json:"name"`
		FathersName string `json:"fathers_name"`
		DateOfBirth string `json:"date_of_birth"`
		PanNumber   string `json:"pan_number"`
	} `json:"onboarding"`
}
type SubmitPanCardAPIResponse struct {
	Onboarding OnBoarding `json:"onboarding"`
}

// read pan card api models
type ReadPanCard struct {
	UserId   string `json:"user_id"`
	ImageUrl string `json:"image_url"`
}

type ReadPanCardAPI struct {
	Onboarding struct {
		ImageUrls []string `json:"image_urls"`
	} `json:"onboarding"`
}

type ReadPanCardAPIResponse struct {
	Name        string     `json:"name"`
	FathersName string     `json:"fathers_name"`
	DateOfBirth string     `json:"date_of_birth"`
	PanNumber   string     `json:"pan_number"`
	Onboarding  OnBoarding `json:"onboarding"`
}

type UploadAadhaarCard struct {
	AadhaarCard *multipart.FileHeader `json:"pan_card"`
	UserId      string                `json:"user_id"`
}

type UploadAadhaarCardAPI struct {
	Onboarding struct {
		AddressProofType string   `json:"address_proof_type"`
		ImageUrls        []string `json:"image_urls"`
	} `json:"onboarding"`
}

type UploadAadhaarCardAPIResponse struct {
	AadhaarUid  string `json:"aadhaar_uid"`
	AadhaarVid  string `json:"aadhaar_vid"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	Pincode     string `json:"pincode"`
	Address     string `json:"address"`
	District    string `json:"district"`
	City        string `json:"city"`
	State       string `json:"state"`
	Gender      string `json:"gender"`
	Onboarding  struct {
		Uuid                 string `json:"uuid"`
		PanNumber            string `json:"pan_number"`
		ExistingInvestor     bool   `json:"existing_investor"`
		Name                 string `json:"name"`
		DateOfBirth          string `json:"date_of_birth"`
		Email                string `json:"email"`
		PhoneNumber          string `json:"phone_number"`
		KycStatus            string `json:"kyc_status"`
		PanCardImageUrl      string `json:"pan_card_image_url"`
		FathersName          string `json:"fathers_name"`
		AddressProofImageUrl string `json:"address_proof_image_url"`
		AddressProofType     string `json:"address_proof_type"`
		Address              string `json:"address"`
		City                 string `json:"city"`
		Pincode              string `json:"pincode"`
		SignatureImageUrl    string `json:"signature_image_url"`
		SelfieImageUrl       string `json:"selfie_image_url"`
		CancelledChequeUrl   string `json:"cancelled_cheque_url"`
		VideoUrl             string `json:"video_url"`
		AnnualIncome         string `json:"annual_income"`
		Gender               string `json:"gender"`
		Occupation           string `json:"occupation"`
		MaritalStatus        string `json:"marital_status"`
		IsFatcaRequired      bool   `json:"is_fatca_required"`
	} `json:"onboarding"`
}

type UploadAadhaarCardResponse struct {
	AadhaarUid  string `json:"aadhaar_uid"`
	AadhaarVid  string `json:"aadhaar_vid"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	Pincode     string `json:"pincode"`
	Address     string `json:"address"`
	District    string `json:"district"`
	City        string `json:"city"`
	State       string `json:"state"`
	Gender      string `json:"gender"`
}

// submit addressProof models
type SubmitAadhaarCardImage struct {
	UserId      string `json:"user_id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	District    string `json:"district"`
	PinCode     string `json:"pin_code"`
	AadharUid   string `json:"aadhar_uid"`
}
type SubmitAadhaarCardImageAPI struct {
	Onboarding struct {
		AddressProofType string `json:"address_proof_type"`
		Name             string `json:"name"`
		ExpiryDate       string `json:"expiry_date"`
		DateOfBirth      string `json:"date_of_birth"`
		IssueDate        string `json:"issue_date"`
		Address          string `json:"address"`
		City             string `json:"city"`
		State            string `json:"state"`
		District         string `json:"district"`
		Pincode          string `json:"pincode"`
		LicenseNumber    string `json:"license_number"`
		AadhaarUid       string `json:"aadhaar_uid"`
		PassportNumber   string `json:"passport_number"`
		VoterIdNumber    string `json:"voter_id_number"`
	} `json:"onboarding"`
}
type SubmitAadhaarCardImageAPIResponse struct {
	Onboarding struct {
		Uuid                 string `json:"uuid"`
		PanNumber            string `json:"pan_number"`
		ExistingInvestor     bool   `json:"existing_investor"`
		Name                 string `json:"name"`
		DateOfBirth          string `json:"date_of_birth"`
		Email                string `json:"email"`
		PhoneNumber          string `json:"phone_number"`
		KycStatus            string `json:"kyc_status"`
		PanCardImageUrl      string `json:"pan_card_image_url"`
		FathersName          string `json:"fathers_name"`
		AddressProofImageUrl string `json:"address_proof_image_url"`
		AddressProofType     string `json:"address_proof_type"`
		Address              string `json:"address"`
		City                 string `json:"city"`
		Pincode              string `json:"pincode"`
		SignatureImageUrl    string `json:"signature_image_url"`
		SelfieImageUrl       string `json:"selfie_image_url"`
		CancelledChequeUrl   string `json:"cancelled_cheque_url"`
		VideoUrl             string `json:"video_url"`
		AnnualIncome         string `json:"annual_income"`
		Gender               string `json:"gender"`
		Occupation           string `json:"occupation"`
		MaritalStatus        string `json:"marital_status"`
		IsFatcaRequired      bool   `json:"is_fatca_required"`
	} `json:"onboarding"`
}

// submitInvestors Details
type SubmitInvestorDetails struct {
	UserId                       string `json:"user_id"`
	Gender                       string `json:"gender" validate:"required,oneof='M','F',T"`
	MaritalStatus                string `json:"marital_status" validate:"required,oneof='MARRIED','UNMARRIED',OTHERS"`
	OccupationDescription        string `json:"occupation_description"`
	OccupationCode               string `json:"occupation_code"`
	CitizenshipCode              string `json:"citizenship_code"`
	CitizenshipCountry           string `json:"citizenship_country"`
	ApplicationStatusCode        string `json:"application_status_code"`
	ApplicationStatusDescription string `json:"application_status_description"`
	AnnualIncome                 string `json:"annual_income"`
}

type SubmitInvestorDetailsAPI struct {
	Onboarding struct {
		Gender                       string `json:"gender"`
		MaritalStatus                string `json:"marital_status"`
		OccupationDescription        string `json:"occupation_description"`
		OccupationCode               string `json:"occupation_code"`
		CitizenshipCode              string `json:"citizenship_code"`
		CitizenshipCountry           string `json:"citizenship_country"`
		ApplicationStatusCode        string `json:"application_status_code"`
		ApplicationStatusDescription string `json:"application_status_description"`
		AnnualIncome                 string `json:"annual_income"`
	} `json:"onboarding"`
}

type SubmitInvestorDetailsAPIResponse struct {
}

// uploadSignature
type UploadSignature struct {
	UserId    string                `json:"user_id"`
	Signature *multipart.FileHeader `json:"signature"`
}

type UploadSignatureAPI struct {
	Onboarding struct {
		ImageUrls []string `json:"image_urls"`
	} `json:"onboarding"`
}
type UploadSignatureAPIResponse struct {
}

// uploadSelfie added
type UploadSelfie struct {
	UserId string                `json:"user_id"`
	Selfie *multipart.FileHeader `json:"selfie"`
}
type UploadSelfieAPI struct {
	Onboarding struct {
		ImageUrls []string `json:"image_urls"`
	} `json:"onboarding"`
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
	UserId string                `json:"user_id"`
	Video  *multipart.FileHeader `json:"video"`
}

type SubmitVideoVerificationAPI struct {
	Onboarding struct {
		VideoUrl      string `json:"video_url"`
		TransactionId string `json:"transaction_id"`
	} `json:"onboarding"`
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

// executeVerification struct
type ExecuteVerification struct {
	UserId string `json:"user_id"`
}
type ExecuteVerificationAPI struct {
	UserId string `json:"user_id"`
}
type ExecuteVerificationAPIResponse struct {
	OnBoarding OnBoarding `json:"on_boarding"`
}

// onBoarding object struct

type OnBoardingObject struct {
	Onboarding OnBoarding `json:"onboarding"`
}
type OnBoarding struct {
	Uuid                 string `json:"uuid"`
	PanNumber            string `json:"pan_number"`
	ExistingInvestor     bool   `json:"existing_investor"`
	Name                 string `json:"name"`
	DateOfBirth          string `json:"date_of_birth"`
	Email                string `json:"email"`
	PhoneNumber          string `json:"phone_number"`
	KycStatus            string `json:"Kyc_status" validate:"required, oneof='success','failure','pending'"`
	PanCardImageUrl      string `json:"pan_card_image_url"`
	FathersName          string `json:"fathers_name"`
	AddressProofImageUrl string `json:"address_proof_image_url"`
	AddressProofType     string `json:"address_proof_type" validate:"required, oneof='aadhar','voter_id','passport','licence'"`
	Address              string `json:"address"`
	City                 string `json:"city"`
	Pincode              string `json:"pincode"`
	SignatureImageUrl    string `json:"signature_image_url"`
	SelfieImageUrl       string `json:"selfie_image_url"`
	CancelledChequeUrl   string `json:"cancelled_cheque_url"`
	VideoUrl             string `json:"video_url"`
	AnnualIncome         string `json:"annual_income"`
	Gender               string `json:"gender"`
	Occupation           string `json:"occupation"`
	MaritalStatus        string `json:"marital_status"`
}
