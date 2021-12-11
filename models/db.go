package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() *gorm.DB {
	dbUri := fmt.Sprintf("host=localhost user=postgres dbname=blogapi sslmode=disable password=postgres")

	conn, err := gorm.Open("postgres", dbUri)

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return conn
}
