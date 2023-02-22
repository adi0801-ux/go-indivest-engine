package main

import (
	"fmt"
	gs "github.com/randree/gormseeder"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=8178954036 dbname=indivest_engine port=5432 sslmode=disable"))
	if err != nil {
		fmt.Println(err.Error())
	}

	gs.InitSeeder(db, "Seeders")
}
