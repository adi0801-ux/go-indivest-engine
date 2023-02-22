package repositories

import (
	"indivest-engine/db"
	"indivest-engine/models"
)

type UserRepository struct {
	Db *db.Database
}

func (s *UserRepository) CreateUserLeads(w *models.UserLeads) error {
	return s.Db.CreateUserLeads_(w)
}

func (s *UserRepository) ReadUserLeads(userId string) (*models.UserLeads, error) {
	return s.Db.ReadUserLeads_(userId)
}

func (s *UserRepository) UpdateOrCreateUserLeads(w *models.UserLeads) error {
	return s.Db.UpdateOrCreateUserLeads_(w)
}
