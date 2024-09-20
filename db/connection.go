package db

import (
	"log"

	"github.com/theinlaoq/booking-api-testcase/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=postgres password=govno dbname=booking port=5432 sslmode=disable"
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database:", err)
	} else {
		DB.AutoMigrate(models.User{}, models.Booking{})
		log.Println("db connected")
	}
}
