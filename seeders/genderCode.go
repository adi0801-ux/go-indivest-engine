package seeders

import (
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) GenderCode() error {
	listOfObject := []models.GenderCodes{
		{Id: 1, Code: "M", Description: "Male"},
		{Id: 2, Code: "F", Description: "Female"},
		{Id: 3, Code: "T", Description: "Transgender"},
	}

	for _, listDb := range listOfObject {
		resp := s.db.Store.FirstOrCreate(&listDb)

		if resp.Error != nil {
			utils.Log.Error(resp.Error)
		}
	}
	return nil
}
