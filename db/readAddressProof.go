package db

import (
	"fmt"
	"indivest-engine/constants"
	"indivest-engine/models"
)

func (d *Database) CreateReadAddressProof_(m *models.ReadAddressProofDB) error {
	result := d.store.Create(&m)
	return result.Error
}

func (d *Database) ReadAddressProof_(userId string) (*models.ReadAddressProofDB, error) {
	u := &models.ReadAddressProofDB{}
	err := d.store.Where("user_id = ?", userId).Find(u).Error
	if u.CreatedAt.String() == constants.StartDateTime {
		return u, fmt.Errorf(constants.UserNotFound)
	}
	return u, err
}
