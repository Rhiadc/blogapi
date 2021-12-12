package config

import (
	"fmt"
	"os"
)

func BuildDSN() string {

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}

	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = "postgres"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "blogapi"
	}

	return fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbname, pass)
}
