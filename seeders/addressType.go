package seeders

import (
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) AddressType() error {
	listOfObject := []models.AddressType{
		{Id: 1, Code: "1", Description: "Residential or Business"},
		{Id: 2, Code: "2", Description: "Residential"},
		{Id: 3, Code: "3", Description: "Business"},
		{Id: 4, Code: "4", Description: "Registered Office"},
	}

	for _, listDb := range listOfObject {
		resp := s.db.Store.FirstOrCreate(&listDb)
		if resp.Error != nil {
			utils.Log.Error(resp.Error)
		}
	}
	return nil
}
