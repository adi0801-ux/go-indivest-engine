package seeders

import (
	"indivest-engine/db"
	"indivest-engine/utils"
)

// //Run Seeder

type Seeder struct {
	db *db.Database
}

func RunSeeders(store *db.Database) error {
	s := &Seeder{db: store}

	err := s.FatcaCountryCode()
	if err != nil {
		utils.Log.Error(err)
	}
	err = s.GenderCode()
	if err != nil {
		utils.Log.Error(err)
	}
	err = s.AddressType()
	if err != nil {
		utils.Log.Error(err)
	}
	err = s.AnnualIncomeCode()
	if err != nil {
		utils.Log.Error(err)
	}
	err = s.ApplicantStatusCode()
	if err != nil {
		utils.Log.Error(err)
	}
	err = s.MaritalStatusCode()
	if err != nil {
		utils.Log.Error(err)
	}
	err = s.SourceOfWealth()
	if err != nil {
		utils.Log.Error(err)
	}
	err = s.OccupationCode()
	if err != nil {
		utils.Log.Error(err)
	}
	err = s.CountryCode()
	if err != nil {
		utils.Log.Error(err)
	}
	return nil

}
