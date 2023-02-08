package repositories

import "indivest-engine/models"

func (s *SavvyRepository) CreateDeposits(w *models.CreateDepositsDb) error {
	return s.Db.CreateDeposits_(w)
}

func (s *SavvyRepository) ReadDeposits(userId string) (*models.CreateDepositsDb, error) {
	return s.Db.ReadDeposits_(userId)

}
func (s *SavvyRepository) CreateSip(w *models.CreateSipDb) error {
	return s.Db.CreateSip_(w)
}
