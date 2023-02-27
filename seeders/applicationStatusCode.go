package seeders

import (
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) ApplicantStatusCode() error {
	listOfObject := []models.ApplicationStatusCode{
		{Id: 1, Code: "R", Description: "Resident Indian"},
		{Id: 2, Code: "N", Description: "Non-Resident Indian"},
		{Id: 3, Code: "P", Description: "Foreign National"},
		{Id: 4, Code: "I", Description: "Person of Indian Origin"},
	}

	for _, listDb := range listOfObject {
		resp := s.db.Store.FirstOrCreate(&listDb)

		if resp.Error != nil {
			utils.Log.Error(resp.Error)
		}
	}
	return nil
}
