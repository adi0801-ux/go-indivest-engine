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

func (s *AddFullKyc) ReadFullKyc(userId string) (*models.StartFullKycDB, error) {
	return s.Db.ReadFullKyc_(userId)
}
