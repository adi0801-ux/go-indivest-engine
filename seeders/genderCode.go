package seeders

import (
	"fmt"
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) GenderCode() error {
	listOfObject := []models.GenderCodes{
		{1, "M", "Male"},
		{2, "F", "Female"},
		{3, "T", "Transgender"},
	}

	for _, listDb := range listOfObject {
		err := s.db.Store.Create(&listDb)
		//if err != nil && err.Error.Error() == {
		//	fmt.Println(err.Error)
		//	utils.Log.Error(err)
		//}
		if err != nil {
			fmt.Println(err)
			utils.Log.Error(err)
		} else {
			return err.Error
		}
	}
	return nil
}
