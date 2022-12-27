package models

import "mime/multipart"

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

type StartFullKycDB struct {
}

type UploadFile struct {
	UserId     string                `json:"user_id"`
	UploadFile *multipart.FileHeader `json:"upload_file"`
	Url        string                `json:"url"`
}

type UploadFileAPI struct {
	Url string `json:"string"`
}
type UploadFileDB struct {
	UserId string `json:"user_id"`
	Url    string `json:"string"`
}

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
