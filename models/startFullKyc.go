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

//
//type ReadAddressProof struct {
//	UserId string `json:"user_id"`
//	AddressProofType constant(
//		home
//		)
//}

// read pan card api models
type ReadPanCard struct {
	UserId   string `json:"user_id"`
	ImageUrl string `json:"image_url"`
}

type ReadPanCardAPI struct {
	ImageUrl string `json:"image_url"`
}

type ReadPanCardAPIResponse struct {
	Onboarding struct {
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
	} `json:"onboarding"`
	Name        string `json:"name"`
	FathersName string `json:"fathers_name"`
	DateOfBirth string `json:"date_of_birth"`
	PanNumber   string `json:"pan_number"`
}
