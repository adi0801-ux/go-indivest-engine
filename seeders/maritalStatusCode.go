package seeders

import (
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) MaritalStatusCode() error {
	listOfObject := []models.MaritalStatusCode{
		{Id: 1, Code: "MARRIED", Description: "Married"},
		{Id: 2, Code: "UNMARRIED", Description: "Unmarried"},
		{Id: 3, Code: "OTHERS", Description: "Others"},
	}

	for _, listDb := range listOfObject {
		resp := s.db.Store.FirstOrCreate(&listDb)

		if resp.Error != nil {
			utils.Log.Error(resp.Error)
		}
	}
	return nil
}
