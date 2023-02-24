package seeders

import (
	"fmt"
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
