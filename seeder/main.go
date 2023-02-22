package main

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	gs "github.com/randree/gormseeder"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: false,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,         // Disable color
		},
	)

	//dns = "host=localhost user=postgres password=postgres dbname=spotwallet port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	gs.InitSeeder(db, "Seeders")
}
