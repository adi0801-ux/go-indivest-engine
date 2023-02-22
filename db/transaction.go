package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (d *Database) CreateDeposits_(m *models.CreateDepositsDb) error {
	result := d.store.Create(&m)
	return result.Error
}

func (d *Database) ReadDeposits_(userId string) (*models.CreateDepositsDb, error) {
	u := &models.CreateDepositsDb{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) ReadAllDeposits_(userId string) (*[]models.CreateDepositsDb, error) {
	u := &[]models.CreateDepositsDb{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	return u, err
}

func (d *Database) ReadAlDeposits_(userId string) ([]models.CreateDepositsDb, error) {
	u := []models.CreateDepositsDb{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	return u, err
}
func (d *Database) ReadDepositsByUUID_(uuid string) (*models.CreateDepositsDb, error) {
	u := &models.CreateDepositsDb{}
	err := d.store.Where("uuid = ?", uuid).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
func (d *Database) CreateOrUpdateDepositUuid_(w *models.CreateDepositsDb) error {
	result := d.store.Model(&w).Where("uuid = ?", w.Uuid).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.store.Create(&w)
		return result.Error
	}

	return result.Error
}

func (d *Database) CreateSip_(m *models.CreateSipDb) error {
	result := d.store.Create(&m)
	return result.Error
}
func (d *Database) ReadSip_(userId string) (*models.CreateSipDb, error) {
	u := &models.CreateSipDb{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
func (d *Database) ReadAllSip_(userId string) (*[]models.CreateSipDb, error) {
	u := &[]models.CreateSipDb{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	return u, err
}
func (d *Database) ReadSipUuid_(uuid string) (*models.CreateSipDb, error) {
	u := &models.CreateSipDb{}
	err := d.store.Where("uuid = ?", uuid).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
func (d *Database) UpdateSip_(w *models.CreateSipDb) error {
	result := d.store.Model(&w).Where("uuid = ?", w.Uuid).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.store.Create(&w)
		return result.Error
	}

	return result.Error
}

func (d *Database) CreateWithdrawal_(m *models.CreateWithdrawalDb) error {
	result := d.store.Create(&m)
	return result.Error
}
func (d *Database) ReadWithdrawal_(withdrwalId string) (*models.CreateWithdrawalDb, error) {
	u := &models.CreateWithdrawalDb{}
	err := d.store.Where("withdrawal_id = ?", withdrwalId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
func (d *Database) ReadWithdrawalUuid_(uuid string) (*models.CreateWithdrawalDb, error) {
	u := &models.CreateWithdrawalDb{}
	err := d.store.Where("uuid = ?", uuid).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
func (d *Database) ReadWithdrawalAll_(userId string) (*models.CreateWithdrawalDb, error) {
	u := &models.CreateWithdrawalDb{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
func (d *Database) ReadAllWithdrawal_(userId string) (*[]models.CreateWithdrawalDb, error) {
	u := &[]models.CreateWithdrawalDb{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	return u, err
}
func (d *Database) UpdateWithdrawal_(w *models.CreateWithdrawalDb) error {
	result := d.store.Model(&w).Where("user_id = ?", w.UserId).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.store.Create(&w)
		return result.Error
	}

	return result.Error
}
func (d *Database) UpdateWithdrawalUuid_(w *models.CreateWithdrawalDb) error {
	result := d.store.Model(&w).Where("uuid = ?", w.Uuid).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.store.Create(&w)
		return result.Error
	}

	return result.Error
}

func (d *Database) CreateWatchList_(m *models.WatchListDb) error {
	result := d.store.Create(&m)
	return result.Error
}
func (d *Database) ReadWatchList_(fundCode string) (*models.WatchListDb, error) {
	u := &models.WatchListDb{}
	err := d.store.Where("fund_code = ?", fundCode).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
func (d *Database) ReadWatchListUserId_(userId string) (*models.WatchListDb, error) {
	u := &models.WatchListDb{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
func (d *Database) DeleteWatchList_(w *models.WatchListDb) error {

	err := d.store.Delete(w).Error
	if err != nil {
		utils.Log.Error(err)
		return err
	}
	return nil
}
