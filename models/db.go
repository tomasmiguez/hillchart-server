package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getDsn() string {
	return os.Getenv("DB_URL")
}

func init() {
	var err error
	DB, err = gorm.Open(postgres.Open(getDsn()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
