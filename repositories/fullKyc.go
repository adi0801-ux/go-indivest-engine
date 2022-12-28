package repositories

import (
	"indivest-engine/db"
	"indivest-engine/models"
)

type AddFullKyc struct {
	Db *db.Database
}

func (s *AddFullKyc) CreateFullKyc(m *models.StartFullKycDB) error {
	return s.Db.CreateFullKyc_(m)
}

//
//func (s *AddFullKyc) ReadFullKyc(userId string) (*models.StartFullKycDB, error) {
//	return s.Db.ReadFullKyc_(userId)
//}

// readPanCard repo
type ReadPanCardRepository struct {
	Db *db.Database
}

func (s *ReadPanCardRepository) CreateReadPanCardDetails(w *models.ReadPanCardDB) error {
	return s.Db.CreateReadPanCardDetails_(w)
}

func (s *ReadPanCardRepository) ReadPanCardDetails(userId string) (*models.ReadPanCardDB, error) {
	return s.Db.ReadPanCardDetails_(userId)
}

// readAddressProof repo
type ReadAddressProofReposotry struct {
	Db *db.Database
}

func (s *ReadAddressProofReposotry) CreateReadAddressProof(w *models.ReadAddressProofDB) error {
	return s.Db.CreateReadAddressProof_(w)
}
func (s *ReadAddressProofReposotry) ReadAddressProof(userId string) (*models.ReadAddressProofDB, error) {
	return s.Db.ReadAddressProof_(userId)
}

// startVideoVerification Repo
type StartVideoVerificationRepository struct {
	Db *db.Database
}

func (s *StartVideoVerificationRepository) CreateVideoVerification(w *models.StartVideoVerificationDB) error {
	return s.Db.CreatVideoVerification_(w)
}
func (s *StartVideoVerificationRepository) ReadVideoVerification(userId string) (*models.StartVideoVerificationDB, error) {
	return s.Db.ReadVideoVerification_(userId)

}

// generateKycContractRepo
type GenerateKycContractRepository struct {
	Db *db.Database
}

func (s *GenerateKycContractRepository) CreateKycContract(w *models.GenerateKycContractDB) error {
	return s.Db.CreateKycContract_(w)
}
func (s *GenerateKycContractRepository) ReadKycContract(userId string) (*models.GenerateKycContractDB, error) {
	return s.Db.ReadKycContract_(userId)

}
