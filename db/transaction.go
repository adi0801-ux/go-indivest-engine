package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
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

func (d *Database) ReadWithdrawalAll_(userId string) (*models.CreateWithdrawalDb, error) {
	u := &models.CreateWithdrawalDb{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) UpdateWithdrawal(w *models.CreateWithdrawalDb) error {
	result := d.store.Model(&w).Where("user_id = ?", w.UserId).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.store.Create(&w)
		return result.Error
	}

	return result.Error
}
