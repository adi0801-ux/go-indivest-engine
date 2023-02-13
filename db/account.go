package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
)

func (d *Database) CreateAccount_(m *models.ShowAccountDB) error {
	result := d.store.Create(&m)
	return result.Error
}

func (d *Database) ReadAccount_(userId string) (*models.ShowAccountDB, error) {
	u := &models.ShowAccountDB{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
func (d *Database) ReadAccountWithAmcId_(userId string, AmcId string) (*models.ShowAccountDB, error) {
	u := &models.ShowAccountDB{}
	err := d.store.Where("user_id = ? and amc_id = ?", userId, AmcId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}

func (d *Database) CreateOrUpdateAccount_(w *models.ShowAccountDB) error {
	result := d.store.Model(&w).Where("user_id = ?", w.UserId).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.store.Create(&w)
		return result.Error
	}

	return result.Error
}
