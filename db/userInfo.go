package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
)

func (d *Database) CreateUserLeads_(w *models.UserLeads) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadUserLeads_(userId string) (*models.UserLeads, error) {
	u := &models.UserLeads{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf("user record not found")
	}
	return u, nil
}
