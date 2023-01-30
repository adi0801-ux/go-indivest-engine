package repositories

import (
	"indivest-engine/db"
	"indivest-engine/models"
)

type AccountRepository struct {
	Db *db.Database
}

func (s *AccountRepository) CreateShowAccount(m *models.ShowAccountDB) error {
	return s.Db.CreateShowAccount_(m)
}
