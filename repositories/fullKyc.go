package repositories

import (
	"indivest-engine/db"
	"indivest-engine/models"
)

type SavvyRepository struct {
	Db *db.Database
}

func (s *SavvyRepository) CreateOnboardingObject(m *models.OnboardingObjectDB) error {
	return s.Db.CreateOnboardingObject_(m)
}

func (s *SavvyRepository) ReadOnboardingObject(userId string) (*models.OnboardingObjectDB, error) {
	return s.Db.ReadOnboardingObject_(userId)
}
func (s *SavvyRepository) ReadOnboardingObjectByUUID(uuid string) (*models.OnboardingObjectDB, error) {
	return s.Db.ReadOnboardingObjectByUUID_(uuid)
}

func (s *SavvyRepository) UpdateOrCreateOnboardingObject(m *models.OnboardingObjectDB) error {
	return s.Db.UpdateOrCreateOnboardingObject_(m)
}

func (s *SavvyRepository) UpdateOrCreateOnboardingObjectUuid(m *models.OnboardingObjectDB) error {
	return s.Db.UpdateOrCreateOnboardingObjectUuid_(m)
}

func (s *SavvyRepository) CreateUserBank(m *models.BankAccountDB) error {
	return s.Db.CreateUserBank_(m)
}

func (s *SavvyRepository) ReadAllOccupationStatus() (*[]models.OccupationDB, error) {
	return s.Db.ReadAllOccupationStatus_()
}

func (s *SavvyRepository) ReadAllGenderCodes() (*[]models.GenderCodesDB, error) {
	return s.Db.ReadAllGenderCodes_()
}

func (s *SavvyRepository) ReadAllMaritalStatusCodes() (*[]models.MaritalStatusCodesDB, error) {
	return s.Db.ReadAllMaritalStatusCodes_()
}

func (s *SavvyRepository) ReadAllCountryCodes() (*[]models.CountryCodesDB, error) {
	return s.Db.ReadAllCountryCodes_()
}

func (s *SavvyRepository) ReadAllAnnualIncomeLevel() (*[]models.IncomeLevelDB, error) {
	return s.Db.ReadAllAnnualIncomeLevel_()
}

func (s *SavvyRepository) CreateFullKyc(m *models.StartFullKycDB) error {
	return s.Db.CreateFullKyc_(m)
}

//
//func (s *AddFullKyc) ReadFullKyc(userId string) (*models.StartFullKycDB, error) {
//	return s.Db.ReadFullKyc_(userId)
//}

func (s *SavvyRepository) CreateReadPanCardDetails(w *models.ReadPanCardDB) error {
	return s.Db.CreateReadPanCardDetails_(w)
}

func (s *SavvyRepository) ReadPanCardDetails(userId string) (*models.ReadPanCardDB, error) {
	return s.Db.ReadPanCardDetails_(userId)
}

func (s *SavvyRepository) CreateReadAddressProof(w *models.ReadAddressProofDB) error {
	return s.Db.CreateReadAddressProof_(w)
}
func (s *SavvyRepository) ReadAddressProof(userId string) (*models.ReadAddressProofDB, error) {
	return s.Db.ReadAddressProof_(userId)
}

func (s *SavvyRepository) CreateVideoVerification(w *models.StartVideoVerificationDB) error {
	return s.Db.CreatVideoVerification_(w)
}
func (s *SavvyRepository) ReadVideoVerification(userId string) (*models.StartVideoVerificationDB, error) {
	return s.Db.ReadVideoVerification_(userId)

}

func (s *SavvyRepository) CreateKycContract(w *models.GenerateKycContractDB) error {
	return s.Db.CreateKycContract_(w)
}
func (s *SavvyRepository) ReadKycContract(userId string) (*models.GenerateKycContractDB, error) {
	return s.Db.ReadKycContract_(userId)

}
