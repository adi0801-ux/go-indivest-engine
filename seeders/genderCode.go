package seeders

import (
	"fmt"
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
		fmt.Println("begining of data input")
		err := s.db.Store.Create(&listDb)
		//if err != nil && err.Error.Error() == {
		//	fmt.Println(err.Error)
		//	utils.Log.Error(err)
		//}
		if err != nil {
			fmt.Println(err)
			utils.Log.Error(err)
		} else {
			fmt.Println("hi there")
		}
	}
	return nil
}
