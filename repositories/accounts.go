package repositories

import (
	"indivest-engine/db"
	"indivest-engine/models"
)

type AccountRepository struct {
	Db *db.Database
}

func (s *AccountRepository) CreateAccount(m *models.ShowAccountDB) error {
	return s.Db.CreateAccount_(m)
}
func (s *AccountRepository) ReadAccount(userId string) (*models.ShowAccountDB, error) {
	return s.Db.ReadAccount_(userId)
}
