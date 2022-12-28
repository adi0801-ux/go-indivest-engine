package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
)

func (d *Database) CreatVideoVerification_(m *models.StartVideoVerificationDB) error {
	result := d.store.Create(&m)
	return result.Error
}

func (d *Database) ReadVideoVerification_(userId string) (*models.StartVideoVerificationDB, error) {
	u := &models.StartVideoVerificationDB{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
