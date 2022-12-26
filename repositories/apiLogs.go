package repositories

import (
	"indivest-engine/db"
	"indivest-engine/models"
)

type ApiLogsRepository struct {
	Db *db.Database
}

func (s *ApiLogsRepository) CreateApiLog(w *models.APILog) error {
	return s.Db.CreateApiLog_(w)
}

func (s *ApiLogsRepository) UpdateApiLog(w *models.APILog) error {
	return s.Db.CreateOrUpdateApiLog_(w)
}
