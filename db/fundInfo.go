package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
)

func (d *Database) CreateOrUpdateFundHousesList_(w *models.FundHousesSupported) error {
	result := d.Store.Model(&w).Where("amc_id = ?", w.AMCID).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.Store.Create(&w)
		return result.Error
	}

	return result.Error
}
func (d *Database) ReadAllFundHousesList_() (*[]models.FundHousesSupported, error) {
	u := &[]models.FundHousesSupported{}
	err := d.Store.Find(u).Error
	return u, err
}

func (d *Database) ReadFundHouseDetailsWithAmcCode_(AMCCode string) (*models.FundHousesSupported, error) {
	u := &models.FundHousesSupported{}
	err := d.Store.Where("amc_code = ?", AMCCode).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.SchemeCodeNotFound)
	}
	return u, err
}

func (d *Database) CreateOrUpdateFundDetails_(w *models.FundsSupported) error {
	result := d.Store.Model(&w).Where("amfi_code = ?", w.AMFICode).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.Store.Create(&w)
		return result.Error
	}

	return result.Error
}
func (d *Database) ReadAllFundDetails_() (*[]models.FundsSupported, error) {
	u := &[]models.FundsSupported{}
	err := d.Store.Find(u).Error
	return u, err
}
func (d *Database) ReadFirstTenFundDetails_() (*[]models.FundsSupported, error) {
	u := &[]models.FundsSupported{}
	err := d.Store.Limit(10).Find(&u).Error
	//err := d.Store.Find(u).Error
	return u, err
}
func (d *Database) ReadFundCategory_() (*[]models.FundCategory, error) {
	u := &[]models.FundCategory{}
	err := d.Store.Model(models.FundsSupported{}).Select("DISTINCT ON (category) category").Find(u).Error
	return u, err
}

func (d *Database) ReadFundDetails_(AMFICode string) (*models.FundsSupported, error) {
	u := &models.FundsSupported{}
	err := d.Store.Where("amfi_code = ?", AMFICode).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.SchemeCodeNotFound)
	}
	return u, err
}

func (d *Database) ReadFundDetailsWithAmcCode_(AMCCode string) (*models.FundsSupported, error) {
	u := &models.FundsSupported{}
	err := d.Store.Where("amc_code = ?", AMCCode).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.SchemeCodeNotFound)
	}
	return u, err
}
