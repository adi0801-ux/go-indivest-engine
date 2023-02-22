package main

import (
	"fmt"
	gm "github.com/randree/gormigrator"
	"gorm.io/gorm"
	"time"
)

func init() {

	gm.Mig(gm.State{

		Tag: "occupationCode",

		Up: func(db *gorm.DB) error {

			type OccupationCode struct {
				ID          int       `gorm:"primary_key"`
				CreateAt    time.Time `gorm:"created_at"`
				Code        string
				Description string
			}

			err := db.AutoMigrate(&OccupationCode{})
			if err != nil {
				fmt.Println(err)
				return err
			}

			return db.Create(&OccupationCode{
				Code:        "01",
				Description: "Private Sector",
			}).Error

		},

		Down: func(db *gorm.DB) error {
			err := db.Migrator().DropTable("occupationCode")

			return err
		},
	})

}
