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
	store *gorm.DB
}

func (d *Database) RunMigrations() (err error) {
	err = d.store.AutoMigrate(&models.SessionManager{},
		&models.UserDetails{},
		&models.UserReports{},
	)
	err = d.store.AutoMigrate(&models.UserWallet{},
		&models.UserMFTransactions{},
		&models.UserMFActiveSIP{},
		&models.UserMFHoldings{},
		&models.UserMFDailyReport{},
	)

	err = d.store.AutoMigrate(
		&models.APILog{},
		&models.OnboardingObjectDB{},
		&models.BankAccountDB{},
		&models.OccupationDB{},
		&models.ReadPanCardDB{},
		&models.StartVideoVerificationDB{},
		&models.ShowAccountDB{},
		&models.CreateDepositsDb{},
		&models.CreateSipDb{},
		&models.CreateWithdrawalDb{},
	)

	err = d.store.AutoMigrate(
		&models.FundsSupported{},
		&models.FundHousesSupported{},
	)

	return err
}

func (d *Database) CloseConnection() (err error) {
	conn, err := d.store.DB()

	if err != nil {
		utils.Log.Error(err)
		return
	}
	err = conn.Close()

	return

}
