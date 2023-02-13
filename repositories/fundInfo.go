package repositories

import (
	"indivest-engine/models"
)

func (s *SavvyRepository) CreateOrUpdateFundHousesList(m *models.FundHousesSupported) error {
	return s.Db.CreateOrUpdateFundHousesList_(m)
}

func (s *SavvyRepository) ReadAllFundHousesList() (*[]models.FundHousesSupported, error) {
	return s.Db.ReadAllFundHousesList_()
}

func (s *SavvyRepository) ReadFundHouseDetailsWithAmcCode(AMCCode string) (*models.FundHousesSupported, error) {
	return s.Db.ReadFundHouseDetailsWithAmcCode_(AMCCode)
}

func (s *SavvyRepository) CreateOrUpdateFundDetails(m *models.FundsSupported) error {
	return s.Db.CreateOrUpdateFundDetails_(m)
}

func (s *SavvyRepository) ReadAllFundDetails() (*[]models.FundsSupported, error) {
	return s.Db.ReadAllFundDetails_()
}

func (s *SavvyRepository) ReadFundDetails(AMFICode string) (*models.FundsSupported, error) {
	return s.Db.ReadFundDetails_(AMFICode)
}

func (s *SavvyRepository) ReadFundDetailsWithAmcCode(AMCCode string) (*models.FundsSupported, error) {
	return s.Db.ReadFundDetailsWithAmcCode_(AMCCode)
}
