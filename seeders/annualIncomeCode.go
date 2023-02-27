package seeders

import (
	"fmt"
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) AnnualIncomeCode() error {
	listOfObject := []models.AnnualIncome{
		{Id: 1, Code: "31", Description: "Below 1 Lac"},
		{Id: 2, Code: "32", Description: "1-5 Lacs"},
		{Id: 3, Code: "33", Description: "5-10 Lacs"},
		{Id: 4, Code: "34", Description: "10-25 Lacs"},
		{Id: 5, Code: "35", Description: "25 Lacs-1 crore"},
		{Id: 6, Code: "36", Description: "1 crore"},
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
