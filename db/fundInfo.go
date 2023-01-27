package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
)

func (d *Database) CreateOrUpdateFundHousesList_(w *models.FundHousesSupported) error {
	result := d.store.Model(&w).Where("amc_id = ?", w.AMCID).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.store.Create(&w)
		return result.Error
	}

	return result.Error
}
func (d *Database) ReadAllFundHousesList_() (*[]models.FundHousesSupported, error) {
	u := &[]models.FundHousesSupported{}
	err := d.store.Find(u).Error
	return u, err
}

func (d *Database) CreateOrUpdateFundDetails_(w *models.FundsSupported) error {
	result := d.store.Model(&w).Where("amfi_code = ?", w.AMFICode).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.store.Create(&w)
		return result.Error
	}

	return result.Error
}
func (d *Database) ReadAllFundDetails_() (*[]models.FundsSupported, error) {
	u := &[]models.FundsSupported{}
	err := d.store.Find(u).Error
	return u, err
}

func (d *Database) ReadFundDetails_(AMFICode string) (*models.FundsSupported, error) {
	u := &models.FundsSupported{}
	err := d.store.Where("amfi_code = ?", AMFICode).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.SchemeCodeNotFound)
	}
	return u, err
}
