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
