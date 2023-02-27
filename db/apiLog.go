package db

import (
	"indivest-engine/models"
)

func (d *Database) CreateApiLog_(w *models.APILog) error {
	result := d.Store.Create(&w)
	return result.Error
}

func (d *Database) CreateOrUpdateApiLog_(
	u *models.APILog) (err error) {
	if d.Store.Where("request_id = ?", u.RequestId).Updates(&u).RowsAffected == 0 {
		err = d.Store.Create(&u).Error
	}
	return nil
}
