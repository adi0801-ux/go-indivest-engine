package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
)

func (d *Database) CreateUser_(w *models.User) error {
	result := d.Store.Create(&w)
	return result.Error
}

func (d *Database) ReadUser_(userId string) (*models.User, error) {
	u := &models.User{}
	err := d.Store.Where("user_id = ?", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf("user record not found")
	}
	return u, nil
}

func (d *Database) ReadUserByEmail_(emailId string) (*models.User, error) {
	u := &models.User{}
	err := d.Store.Where("email_id = ?", emailId).Find(u).Error
	if err != nil {
		return u, err
	}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf("user record not found")
	}
	return u, nil
}

func (d *Database) UpdateOrCreateUser_(w *models.User) error {
	result := d.Store.Model(&w).Where("user_id = ?", w.UserId).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.Store.Create(&w)
		return result.Error
	}

	return result.Error
}

func (d *Database) CreateUserLeads_(w *models.UserLeads) error {
	result := d.Store.Create(&w)
	return result.Error
}

func (d *Database) ReadUserLeads_(userId string) (*models.UserLeads, error) {
	u := &models.UserLeads{}
	err := d.Store.Where("user_id = ?", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if u.CreatedAt.String() == constants.StartDateTime {

		return u, fmt.Errorf("user record not found")
	}
	return u, nil
}

func (d *Database) UpdateOrCreateUserLeads_(w *models.UserLeads) error {
	result := d.Store.Model(&w).Where("user_id = ?", w.UserId).Updates(&w)
	if result.RowsAffected == 0 {
		result = d.Store.Create(&w)
		return result.Error
	}

	return result.Error
}
