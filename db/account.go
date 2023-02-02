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
