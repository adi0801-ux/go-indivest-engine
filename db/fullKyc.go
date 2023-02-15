package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
)

func (d *Database) CreateOnboardingObject_(m *models.OnboardingObjectDB) error {
	result := d.store.Create(&m)
	return result.Error
}

func (d *Database) ReadOnboardingObject_(userId string) (*models.OnboardingObjectDB, error) {
	u := &models.OnboardingObjectDB{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) ReadOnboardingObjectByUUID_(uuid string) (*models.OnboardingObjectDB, error) {
	u := &models.OnboardingObjectDB{}
	err := d.store.Where("uuid = ?", uuid).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) UpdateOrCreateOnboardingObject_(w *models.OnboardingObjectDB) error {
	result := d.store.Model(&w).Where("user_id = ?", w.UserId).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.store.Create(&w)
		return result.Error
	}

	return result.Error
}

func (d *Database) UpdateOrCreateOnboardingObjectUuid_(w *models.OnboardingObjectDB) error {
	result := d.store.Model(&w).Where("uuid = ?", w.Uuid).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.store.Create(&w)
		return result.Error
	}

	return result.Error
}
func (d *Database) CreateUserBank_(m *models.BankAccountDB) error {
	result := d.store.Create(&m)
	return result.Error
}

func (d *Database) ReadAllOccupationStatus_() (*[]models.OccupationDB, error) {
	u := &[]models.OccupationDB{}
	err := d.store.Find(u).Error
	return u, err
}

func (d *Database) ReadAllGenderCodes_() (*[]models.GenderCodesDB, error) {
	u := &[]models.GenderCodesDB{}
	err := d.store.Find(u).Error
	return u, err
}

func (d *Database) ReadAllMaritalStatusCodes_() (*[]models.MaritalStatusCodesDB, error) {
	u := &[]models.MaritalStatusCodesDB{}
	err := d.store.Find(u).Error
	return u, err
}

func (d *Database) ReadAllCountryCodes_() (*[]models.CountryCodesDB, error) {
	u := &[]models.CountryCodesDB{}
	err := d.store.Find(u).Error
	return u, err
}

func (d *Database) ReadAllAnnualIncomeLevel_() (*[]models.IncomeLevelDB, error) {
	u := &[]models.IncomeLevelDB{}
	err := d.store.Find(u).Error
	return u, err
}

func (d *Database) CreateFullKyc_(m *models.StartFullKycDB) error {
	result := d.store.Create(&m)
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
	result := d.store.Create(&m)
	return result.Error
}

func (d *Database) ReadPanCardDetails_(userId string) (*models.ReadPanCardDB, error) {
	u := &models.ReadPanCardDB{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) CreateReadAddressProof_(m *models.ReadAddressProofDB) error {
	result := d.store.Create(&m)
	return result.Error
}

func (d *Database) ReadAddressProof_(userId string) (*models.ReadAddressProofDB, error) {
	u := &models.ReadAddressProofDB{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) CreateKycContract_(m *models.GenerateKycContractDB) error {
	result := d.store.Create(&m)
	return result.Error
}

func (d *Database) CreatVideoVerification_(m *models.StartVideoVerificationDB) error {
	result := d.store.Create(&m)
	return result.Error
}

func (d *Database) ReadVideoVerification_(userId string) (*models.StartVideoVerificationDB, error) {
	u := &models.StartVideoVerificationDB{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) ReadKycContract_(userId string) (*models.GenerateKycContractDB, error) {
	u := &models.GenerateKycContractDB{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
