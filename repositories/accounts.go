package repositories

import (
	"indivest-engine/models"
)

//type AccountRepository struct {
//	Db *db.Database
//}

func (s *SavvyRepository) CreateAccount(m *models.ShowAccountDB) error {
	return s.Db.CreateAccount_(m)
}
func (s *SavvyRepository) ReadAccount(userId string) (*models.ShowAccountDB, error) {
	return s.Db.ReadAccount_(userId)
}

func (s *SavvyRepository) CreateOrUpdateAccount(m *models.ShowAccountDB) error {
	return s.Db.CreateOrUpdateAccount_(m)
}
func (s *SavvyRepository) ReadAccountWithAmcId(userId string, AmcId string) (*models.ShowAccountDB, error) {
	return s.Db.ReadAccountWithAmcId_(userId, AmcId)
}
