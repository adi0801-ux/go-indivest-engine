package repositories

import "indivest-engine/models"

func (s *SavvyRepository) CreateDeposits(w *models.CreateDepositsDb) error {
	return s.Db.CreateDeposits_(w)
}

func (s *SavvyRepository) ReadDeposits(userId string) (*models.CreateDepositsDb, error) {
	return s.Db.ReadDeposits_(userId)
}

func (s *SavvyRepository) ReadDepositsByUUID(uuid string) (*models.CreateDepositsDb, error) {
	return s.Db.ReadDepositsByUUID_(uuid)
}
func (s *SavvyRepository) CreateOrUpdateDeposit(m *models.CreateDepositsDb) error {
	return s.Db.CreateOrUpdateDeposit_(m)
}

func (s *SavvyRepository) CreateSip(w *models.CreateSipDb) error {
	return s.Db.CreateSip_(w)
}
func (s *SavvyRepository) ReadSip(userId string) (*models.CreateSipDb, error) {
	return s.Db.ReadSip_(userId)
}
func (s *SavvyRepository) CreateWithdrawal(w *models.CreateWithdrawalDb) error {
	return s.Db.CreateWithdrawal_(w)
}

func (s *SavvyRepository) ReadWithdrawal(withdrwalId string) (*models.CreateWithdrawalDb, error) {
	return s.Db.ReadWithdrawal_(withdrwalId)

}
func (s *SavvyRepository) ReadWithdrawalAll(userId string) (*models.CreateWithdrawalDb, error) {
	return s.Db.ReadWithdrawalAll_(userId)
}
func (s *SavvyRepository) UpdateWithdrawal(m *models.CreateWithdrawalDb) error {
	return s.Db.UpdateWithdrawal(m)
}
