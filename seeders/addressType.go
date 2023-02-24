package seeders

import (
	"fmt"
	"indivest-engine/models"
	"indivest-engine/utils"
)

func (s *Seeder) AddressType() error {
	listOfObject := []models.AddressType{
		{Id: 1, Code: "1", Description: "Residential or Business"},
		{Id: 2, Code: "2", Description: "Residential"},
		{Id: 3, Code: "3", Description: "Business"},
		{Id: 4, Code: "4", Description: "Registered Office"},
	}

	for _, listDb := range listOfObject {
		resp := s.db.Store.Create(&listDb)
		if resp != nil && resp.Error.Error() == "address_types_pkey" {
			//fmt.Println(resp.Error)
			utils.Log.Error(resp.Error)
		} else {
			utils.Log.Error(resp.Error)
		}
		fmt.Println(resp.Error, resp.RowsAffected)
		//if err != nil {
		//	fmt.Println(err.Error)
		//	utils.Log.Error(err)
		//} else {
		//	return err.Error
		//}
	}
	return nil
}
