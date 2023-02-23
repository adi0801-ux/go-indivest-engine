package seeders

import (
	"fmt"
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) SourceOfWealth() error {
	listOfObject := []models.SourceOfWealth{
		{1, "1", "Salary"},
		{2, "2", "Business"},
		{3, "3", "Gift"},
		{4, "4", "Ancestral Property"},
		{5, "5", "Rental Income"},
		{6, "6", "Prize Money"},
		{7, "7", "Royalty"},
		{8, "8", "Others"},
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
