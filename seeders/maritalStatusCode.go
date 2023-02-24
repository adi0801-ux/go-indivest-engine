package seeders

import (
	"fmt"
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
