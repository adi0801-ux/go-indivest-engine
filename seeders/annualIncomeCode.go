package seeders

import (
	"fmt"
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) AnnualIncomeCode() error {
	listOfObject := []models.AnnualIncome{
		{1, "31", "Below 1 Lac"},
		{2, "32", "1-5 Lacs"},
		{3, "33", "5-10 Lacs"},
		{4, "34", "10-25 Lacs"},
		{5, "35", "25 Lacs-1 crore"},
		{6, "36", "1 crore"},
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
