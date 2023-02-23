package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
)

func (d *Database) CreateOnboardingObject_(m *models.OnboardingObjectDB) error {
	result := d.Store.Create(&m)
	return result.Error
}

func (d *Database) ReadOnboardingObject_(userId string) (*models.OnboardingObjectDB, error) {
	u := &models.OnboardingObjectDB{}
	err := d.Store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) ReadOnboardingObjectByUUID_(uuid string) (*models.OnboardingObjectDB, error) {
	u := &models.OnboardingObjectDB{}
	err := d.Store.Where("uuid = ?", uuid).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) UpdateOrCreateOnboardingObject_(w *models.OnboardingObjectDB) error {
	result := d.Store.Model(&w).Where("user_id = ?", w.UserId).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.Store.Create(&w)
		return result.Error
	}

	return result.Error
}

func (d *Database) UpdateOrCreateOnboardingObjectUuid_(w *models.OnboardingObjectDB) error {
	result := d.Store.Model(&w).Where("uuid = ?", w.Uuid).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.Store.Create(&w)
		return result.Error
	}

	return result.Error
}
func (d *Database) CreateUserBank_(m *models.BankAccountDB) error {
	result := d.Store.Create(&m)
	return result.Error
}

func (d *Database) ReadAllOccupationStatus_() (*[]models.OccupationCode, error) {
	u := &[]models.OccupationCode{}
	err := d.Store.Find(u).Error
	return u, err
}

func (d *Database) ReadAllGenderCodes_() (*[]models.GenderCodes, error) {
	u := &[]models.GenderCodes{}
	err := d.Store.Find(u).Error
	return u, err
}

func (d *Database) ReadAllMaritalStatusCodes_() (*[]models.MaritalStatusCode, error) {
	u := &[]models.MaritalStatusCode{}
	err := d.Store.Find(u).Error
	return u, err
}

func (d *Database) ReadAllCountryCodes_() (*[]models.CountryCode, error) {
	u := &[]models.CountryCode{}
	err := d.Store.Find(u).Error
	return u, err
}

func (d *Database) ReadAllAnnualIncomeLevel_() (*[]models.AnnualIncome, error) {
	u := &[]models.AnnualIncome{}
	err := d.Store.Find(u).Error
	return u, err
}
func (d *Database) ReadSourceOfWealth_() (*[]models.SourceOfWealth, error) {
	u := &[]models.SourceOfWealth{}
	err := d.Store.Find(u).Error
	return u, err
}
func (d *Database) ReadFatcaCountryCode() (*[]models.FatcaCountryCode, error) {
	u := &[]models.FatcaCountryCode{}
	err := d.Store.Find(u).Error
	return u, err
}
func (d *Database) ReadApplicationStatus() (*[]models.ApplicationStatusCode, error) {
	u := &[]models.ApplicationStatusCode{}
	err := d.Store.Find(u).Error
	return u, err
}
func (d *Database) ReadAddressType() (*[]models.AddressType, error) {
	u := &[]models.AddressType{}
	err := d.Store.Find(u).Error
	return u, err
}

func (d *Database) CreateFullKyc_(m *models.StartFullKycDB) error {
	result := d.Store.Create(&m)
	return result.Error
}

//func (d *Database) ReadFullKyc_(userId string) (*models.StartFullKycDB, error) {
//	u := &models.StartFullKycDB{}
//	err := d.store.Where("user_id = ?", userId).Find(u).Error
//	if u.CreatedAt.String() == constants.StartDateTime {
//		return u, fmt.Errorf(constants.UserNotFound)
//	}
//	return u, err
//}

func (d *Database) CreateReadPanCardDetails_(m *models.ReadPanCardDB) error {
	result := d.Store.Create(&m)
	return result.Error
}

func (d *Database) ReadPanCardDetails_(userId string) (*models.ReadPanCardDB, error) {
	u := &models.ReadPanCardDB{}
	err := d.Store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) CreateReadAddressProof_(m *models.ReadAddressProofDB) error {
	result := d.Store.Create(&m)
	return result.Error
}

func (d *Database) ReadAddressProof_(userId string) (*models.ReadAddressProofDB, error) {
	u := &models.ReadAddressProofDB{}
	err := d.Store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) CreateKycContract_(m *models.GenerateKycContractDB) error {
	result := d.Store.Create(&m)
	return result.Error
}

func (d *Database) CreatVideoVerification_(m *models.StartVideoVerificationDB) error {
	result := d.Store.Create(&m)
	return result.Error
}

func (d *Database) ReadVideoVerification_(userId string) (*models.StartVideoVerificationDB, error) {
	u := &models.StartVideoVerificationDB{}
	err := d.Store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) ReadKycContract_(userId string) (*models.GenerateKycContractDB, error) {
	u := &models.GenerateKycContractDB{}
	err := d.Store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
