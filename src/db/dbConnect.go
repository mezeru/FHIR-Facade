package db

import (
	"log"

	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect2Pg() {
	dsn := "host=localhost user=fhir password=fhir dbname=fhir port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to db")
	}

	db.AutoMigrate(&Patient{})

}
