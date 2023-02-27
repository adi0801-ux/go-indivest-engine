package seeders

import (
	"indivest-engine/db"
)

// //Run Seeder

type Seeder struct {
	db *db.Database
}

func RunAddressType(store *db.Database) error {
	s := &Seeder{db: store}
	err := s.AddressType()
	if err != nil {
		return err
	}
	return err
}

func RunFatcaCountryCode(store *db.Database) error {
	s := &Seeder{db: store}
	err := s.FatcaCountryCode()
	if err != nil {
		return err
	}
	return err
}
func RunAnnualIncomeCode(store *db.Database) error {
	s := &Seeder{db: store}
	err := s.AnnualIncomeCode()
	//err := s.SourceOfWealth()
	if err != nil {
		return err
	}
	return err
}
func RunApplicationStatusCode(store *db.Database) error {
	s := &Seeder{db: store}
	err := s.ApplicantStatusCode()
	//err := s.SourceOfWealth()
	if err != nil {
		return err
	}
	return err
}

func RunCountryCode(store *db.Database) error {
	s := &Seeder{db: store}
	err := s.CountryCode()
	//err := s.SourceOfWealth()
	if err != nil {
		return err
	}
	return err
}
func RunMaritalStatus(store *db.Database) error {
	s := &Seeder{db: store}
	err := s.MaritalStatusCode()
	//err := s.SourceOfWealth()
	if err != nil {
		return err
	}
	return err
}
func RunOccupationCodes(store *db.Database) error {
	s := &Seeder{db: store}
	err := s.OccupationCode()
	//err := s.SourceOfWealth()
	if err != nil {
		return err
	}
	return err
}
func RunSourceOfWealthCodes(store *db.Database) error {
	s := &Seeder{db: store}
	err := s.SourceOfWealth()
	if err != nil {
		return err
	}
	return err
}
func RunGenderCodes(store *db.Database) error {
	s := &Seeder{db: store}
	err := s.GenderCode()
	//err := s.SourceOfWealth()
	if err != nil {
		return err
	}
	return err
}
