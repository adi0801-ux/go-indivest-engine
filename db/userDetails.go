package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
)

func (d *Database) CreateUserDetails_(w *models.UserDetails) error {
	result := d.store.Create(&w)
	return result.Error
}

func (d *Database) ReadUserDetails_(userId string) (*models.UserDetails, error) {
	u := &models.UserDetails{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if err != nil {
		return u, err
	}
	if u.CreatedAt.String()== constants.StartDateTime{

		return u , fmt.Errorf("user record not found")
	}
	return u, nil
}

func (d *Database) UpdateUserDetails_(userDetails *models.UserDetails) error {

	result := d.store.Where("user_id = ?", userDetails.UserID).Updates(userDetails)

	return result.Error
}

func (d *Database) UpdateUserDetailWithField_(userId string , fieldName string , value interface{}) error {

	result := d.store.Where("user_id = ?", userId).Update(fieldName,value)

	return result.Error
}

func (d *Database) DeleteUserDetails_(userID string) error {
	userDetails := &models.UserDetails{}
	err := d.store.Where("user_id = ?", userID).First(userDetails).Error
	if err != nil {
		return err
	}

	err = d.store.Delete(userDetails, userID).Error
	if err != nil {
		return err
	}
	return nil
}


func (d *Database) CreateOrUpdateUserDetails_(
	userDetails *models.UserDetails) (err error) {

	if d.store.Where("user_id = ?", userDetails.UserID).Updates(&userDetails).RowsAffected == 0 {
		err = d.store.Create(&userDetails).Error
	}
	return nil
}