package seeders

import (
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) SourceOfWealth() error {
	listOfObject := []models.SourceOfWealth{
		{Id: 1, Code: "1", Description: "Salary"},
		{Id: 2, Code: "2", Description: "Business"},
		{Id: 3, Code: "3", Description: "Gift"},
		{Id: 4, Code: "4", Description: "Ancestral Property"},
		{Id: 5, Code: "5", Description: "Rental Income"},
		{Id: 6, Code: "6", Description: "Prize Money"},
		{Id: 7, Code: "7", Description: "Royalty"},
		{Id: 8, Code: "8", Description: "Others"},
	}

	for _, listDb := range listOfObject {
		resp := s.db.Store.FirstOrCreate(&listDb)

		if resp.Error != nil {
			utils.Log.Error(resp.Error)
		}
	}
	return nil
}
