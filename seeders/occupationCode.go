package seeders

import (
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) OccupationCode() error {
	listOfObject := []models.OccupationCode{
		{Id: 1, Code: "1", Description: "Private Sector"},
		{Id: 2, Code: "2", Description: "Public Sector"},
		{Id: 3, Code: "3", Description: "Business"},
		{Id: 4, Code: "4", Description: "Professional"},
		{Id: 5, Code: "6", Description: "Retired"},
		{Id: 6, Code: "7", Description: "Housewife"},
		{Id: 7, Code: "8", Description: "Student"},
		{Id: 8, Code: "10", Description: "Government Sector"},
		{Id: 9, Code: "99", Description: "Others"},
		{Id: 10, Code: "11", Description: "Self Employed"},
		{Id: 11, Code: "12", Description: "Not Categorized"},
	}

	for _, listDb := range listOfObject {
		resp := s.db.Store.FirstOrCreate(&listDb)

		if resp.Error != nil {
			utils.Log.Error(resp.Error)
		}
	}
	return nil
}
