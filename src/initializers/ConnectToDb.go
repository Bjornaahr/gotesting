package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	dbHost := os.Getenv("dbhost")
	dbUser := os.Getenv("dbusername")
	dbPass := os.Getenv("dbpass")
	dbName := os.Getenv("dbname")
	dbPort := os.Getenv("dbport")
	dbSSL := os.Getenv("dbssl")

	caCertString := "ca-certificate.crt"

	var err error

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s sslrootcert='%s'", dbHost, dbUser, dbPass, dbName, dbPort, dbSSL, caCertString),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
