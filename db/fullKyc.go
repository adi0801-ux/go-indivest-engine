package db

import (
	"indivest-engine/models"
)

func (d *Database) CreateFullKyc_(m *models.StartFullKycDB) error {
	result := d.store.Create(&m)
	return result.Error
}

//func (d *Database) ReadFullKyc_(userId string) (*models.StartFullKycDB, error) {
//	u := &models.StartFullKycDB{}
//	err := d.store.Where("user_id = ?", userId).Find(u).Error
//	if u.CreatedAt.String() == constants.StartDateTime {
//		return u, fmt.Errorf(constants.UserNotFound)
//	}
//	return u, err
//}
