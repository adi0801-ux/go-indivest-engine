package seeders

import (
	"fmt"
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) OccupationCode() error {
	listOfObject := []models.OccupationCode{
		{1, "1", "Private Sector"},
		{2, "2", "Public Sector"},
		{3, "3", "Business"},
		{4, "4", "Professional"},
		{5, "6", "Retired"},
		{6, "7", "Housewife"},
		{7, "8", "Student"},
		{8, "10", "Government Sector"},
		{9, "99", "Others"},
		{10, "11", "Self Employed"},
		{11, "12", "Not Categorized"},
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
