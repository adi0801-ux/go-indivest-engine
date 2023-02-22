package repositories

import (
	"indivest-engine/db"
	"indivest-engine/models"
)

type UserDetailsRepository struct {
	Db *db.Database
}

func (u *UserRepository) CreateUserDetails(userDetails *models.UserDetails) error {
	err := u.Db.CreateUserDetails_(userDetails)
	return err
}

func (u *UserRepository) ReadUserDetails(userId string) (*models.UserDetails, error) {
	userDetails, err := u.Db.ReadUserDetails_(userId)

	return userDetails, err

}

func (u *UserRepository) UpdateUserDetails(userDetails *models.UserDetails) error {
	err := u.Db.UpdateUserDetails_(userDetails)
	return err
}

func (u *UserRepository) UpdateUserDetailWithField(userId string, fieldName string, value interface{}) error {
	err := u.Db.UpdateUserDetailWithField_(userId, fieldName, value)
	return err
}

func (u *UserRepository) DeleteUserDetails_(userId string) error {
	err := u.Db.DeleteUserDetails_(userId)
	return err
}

func (u *UserRepository) CreateOrUpdateUserDetails(userDetails *models.UserDetails) error {

	return u.Db.CreateOrUpdateUserDetails_(userDetails)
}
