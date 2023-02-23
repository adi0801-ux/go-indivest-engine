package seeders

import (
	"fmt"
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) ApplicantStatusCode() error {
	listOfObject := []models.ApplicationStatusCode{
		{1, "R", "Resident Indian"},
		{2, "N", "Non-Resident Indian"},
		{3, "P", "Foreign National"},
		{4, "I", "Person of Indian Origin"},
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
