package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"indivest-engine/utils"
	"log"
	"os"
	"time"
)

func ConnectToDB(config *ConnectionConfig) (*Database, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	//dns = "host=localhost user=postgres password=postgres dbname=spotwallet port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		utils.Log.Error(err)
		return &Database{}, err
	}

	database := &Database{
		store: db,
	}

	//err = gormdb.AutoMigrate(&MarginRecords{} , &CancelOrders{}, &TradingPairs{})
	if err != nil {
		return database, err
	}

	return database, nil
}
