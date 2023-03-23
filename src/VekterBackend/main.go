package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		dbHost := os.Getenv("dbhost")
		dbUser := os.Getenv("dbusername")
		dbPass := os.Getenv("dbpass")
		dbName := os.Getenv("dbname")
		dbPort := os.Getenv("dbport")
		dbSSL := os.Getenv("dbssl")

		caCertString := "ca-certificate.crt"

		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  fmt.Sprintf("host=%s user=%s password=%s dbname=%s port='%s', sslmode=%s, sslrootcert=%s", dbHost, dbUser, dbPass, dbName, dbPort, dbSSL, caCertString),
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Migrate the schema
		db.AutoMigrate(&Product{})
		db.AutoMigrate(&User{})

		// Create
		db.Create(&Product{Code: "D42", Price: 100})

		user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
		result := db.Create(&user) // pass pointer of data to Create

		fmt.Printf(strconv.FormatInt(result.RowsAffected, 10))

		// Read
		var product Product
		db.First(&product, 8) // find product with integer primary key
		//db.First(&product, "code = ?", "D42") // find product with code D42

		// Update - update product's price to 200
		db.Model(&product).Update("Price", 2000)

		c.JSONP(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}
