package db

import "indivest-engine/models"

func (d *Database) CreateShowAccount_(m *models.ShowAccountDB) error {
	result := d.store.Create(&m)
	return result.Error
}
