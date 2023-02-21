package repositories

import "indivest-engine/models"

func (s *SavvyRepository) CreateDeposits(w *models.CreateDepositsDb) error {
	return s.Db.CreateDeposits_(w)
}

func (s *SavvyRepository) ReadDeposits(userId string) (*models.CreateDepositsDb, error) {
	return s.Db.ReadDeposits_(userId)
}
func (s *SavvyRepository) ReadAllDeposits(userId string) (*[]models.CreateDepositsDb, error) {
	return s.Db.ReadAllDeposits_(userId)
}

func (s *SavvyRepository) ReadDepositsByUUID(uuid string) (*models.CreateDepositsDb, error) {
	return s.Db.ReadDepositsByUUID_(uuid)
}
func (s *SavvyRepository) CreateOrUpdateDepositUuid(m *models.CreateDepositsDb) error {
	return s.Db.CreateOrUpdateDepositUuid_(m)
}

func (s *SavvyRepository) CreateSip(w *models.CreateSipDb) error {
	return s.Db.CreateSip_(w)
}
func (s *SavvyRepository) ReadSip(userId string) (*models.CreateSipDb, error) {
	return s.Db.ReadSip_(userId)
}
func (s *SavvyRepository) ReadAllSip(userId string) (*[]models.CreateSipDb, error) {
	return s.Db.ReadAllSip_(userId)
}
func (s *SavvyRepository) ReadSipUuid(uuid string) (*models.CreateSipDb, error) {
	return s.Db.ReadSipUuid_(uuid)
}
func (s *SavvyRepository) UpdateSip(m *models.CreateSipDb) error {
	return s.Db.UpdateSip_(m)
}

func (s *SavvyRepository) CreateWithdrawal(w *models.CreateWithdrawalDb) error {
	return s.Db.CreateWithdrawal_(w)
}

func (s *SavvyRepository) ReadWithdrawal(withdrwalId string) (*models.CreateWithdrawalDb, error) {
	return s.Db.ReadWithdrawal_(withdrwalId)

}
func (s *SavvyRepository) ReadWithdrawalUuid(uuid string) (*models.CreateWithdrawalDb, error) {
	return s.Db.ReadWithdrawalUuid_(uuid)

}
func (s *SavvyRepository) ReadWithdrawalAll(userId string) (*models.CreateWithdrawalDb, error) {
	return s.Db.ReadWithdrawalAll_(userId)
}

func (s *SavvyRepository) ReadAllWithdrawal(userId string) (*[]models.CreateWithdrawalDb, error) {
	return s.Db.ReadAllWithdrawal_(userId)
}
func (s *SavvyRepository) UpdateWithdrawal(m *models.CreateWithdrawalDb) error {
	return s.Db.UpdateWithdrawal_(m)
}
func (s *SavvyRepository) UpdateWithdrawalUuid(m *models.CreateWithdrawalDb) error {
	return s.Db.UpdateWithdrawalUuid_(m)
}

func (s *SavvyRepository) CreateWatchList(w *models.WatchListDb) error {
	return s.Db.CreateWatchList_(w)
}

func (s *SavvyRepository) ReadWatchList(fundCode string) (*models.WatchListDb, error) {
	return s.Db.ReadWatchList_(fundCode)
}

func (s *SavvyRepository) ReadWatchListUserId(UserId string) (*models.WatchListDb, error) {
	return s.Db.ReadWatchListUserId_(UserId)
}

func (s *SavvyRepository) DeleteWatchList(w *models.WatchListDb) (*models.WatchListDb, error) {
	err := s.Db.DeleteWatchList_(w)
	return nil, err
}
