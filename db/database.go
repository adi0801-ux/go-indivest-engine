package db

import (
	"gorm.io/gorm"
	"indivest-engine/models"
	"indivest-engine/utils"
)

type ConnectionConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
	DSN      string
}

type Database struct {
	Store *gorm.DB
}

func (d *Database) RunMigrations() (err error) {
	err = d.Store.AutoMigrate(&models.SessionManager{},
		&models.UserDetails{},
		&models.UserReports{},
	)
	err = d.Store.AutoMigrate(&models.UserWallet{},
		&models.UserMFTransactions{},
		&models.UserMFActiveSIP{},
		&models.UserMFHoldings{},
		&models.UserMFDailyReport{},
	)

	err = d.Store.AutoMigrate(
		&models.APILog{},
		&models.OnboardingObjectDB{},
		&models.BankAccountDB{},
		&models.ReadPanCardDB{},
		&models.StartVideoVerificationDB{},
		&models.ShowAccountDB{},
		&models.CreateDepositsDb{},
		&models.CreateSipDb{},
		&models.CreateWithdrawalDb{},
	)

	err = d.Store.AutoMigrate(
		&models.FundsSupported{},
		&models.FundHousesSupported{},
		&models.WatchListDb{},
	)

	err = d.Store.AutoMigrate(
		&models.UserLeads{},
		&models.User{},
		&models.FatcaCountryCode{},
		&models.CountryCode{},
		&models.AddressType{},
		&models.AnnualIncome{},
		&models.ApplicationStatusCode{},
		&models.MaritalStatusCode{},
		&models.SourceOfWealth{},
		&models.OccupationCode{},
		&models.GenderCodes{},
	)

	return err
}

func (d *Database) CloseConnection() (err error) {
	conn, err := d.Store.DB()

	if err != nil {
		utils.Log.Error(err)
		return
	}
	err = conn.Close()

	return

}
