package seeders

import (
	"fmt"
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) MaritalStatusCode() error {
	listOfObject := []models.MaritalStatusCode{
		{1, "MARRIED", "Married"},
		{2, "UNMARRIED", "Unmarried"},
		{3, "OTHERS", "Others"},
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
