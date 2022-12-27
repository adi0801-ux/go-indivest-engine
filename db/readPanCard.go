package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
)

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
