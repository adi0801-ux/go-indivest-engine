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
