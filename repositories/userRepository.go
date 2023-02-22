package repositories

import (
	"indivest-engine/db"
	"indivest-engine/models"
)

type UserRepository struct {
	Db *db.Database
}

func (s *UserRepository) CreateUser(w *models.User) error {
	return s.Db.CreateUser_(w)
}

func (s *UserRepository) ReadUser(userId string) (*models.User, error) {
	return s.Db.ReadUser_(userId)
}

func (s *UserRepository) ReadUserByEmail(emailId string) (*models.User, error) {
	return s.Db.ReadUserByEmail_(emailId)
}

func (s *UserRepository) UpdateOrCreateUser(w *models.User) error {
	return s.Db.UpdateOrCreateUser_(w)
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
