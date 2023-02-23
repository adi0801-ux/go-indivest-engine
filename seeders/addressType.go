package seeders

import (
	"fmt"
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) AddressType() error {
	listOfObject := []models.AddressType{
		{1, "1", "Residential or Business"},
		{2, "2", "Residential"},
		{3, "3", "Business"},
		{4, "4", "Registered Office"},
	}

	for _, listDb := range listOfObject {
		err := s.db.Store.Create(&listDb)
		//if err != nil && err.Error.Error() == {
		//	fmt.Println(err.Error)
		//	utils.Log.Error(err)
		//}
		if err != nil {
			fmt.Println(err.Error)
			utils.Log.Error(err)
		} else {
			return err.Error
		}
	}
	return nil
}
