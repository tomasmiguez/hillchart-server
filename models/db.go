package models

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/tomasmiguez/hillchart-server/configs"
)

var DB *gorm.DB

func getDsn() string {
	dbConfig := configs.Db()
	return fmt.Sprintf("host=%s user=%s dbname=%s port=%s", dbConfig.Host, dbConfig.User, dbConfig.DbName, dbConfig.Port)
}

func init() {
    var err error
	DB, err = gorm.Open(postgres.Open(getDsn()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
