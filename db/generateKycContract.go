package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
)

func (d *Database) CreateKycContract_(m *models.GenerateKycContractDB) error {
	result := d.store.Create(&m)
	return result.Error
}

func (d *Database) ReadKycContract_(userId string) (*models.GenerateKycContractDB, error) {
	u := &models.GenerateKycContractDB{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
