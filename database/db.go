package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rhiadc/blogapi/database/config"
)

func Connect() *gorm.DB {

	conn, err := gorm.Open("postgres", config.BuildDSN())

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return conn
}
