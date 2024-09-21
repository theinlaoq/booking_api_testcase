package db

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"github.com/theinlaoq/booking-api-testcase/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialzing config: %s", err.Error())
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		viper.GetString("host"), viper.GetString("user"), viper.GetString("password"),
		viper.GetString("dbname"), viper.GetString("dbport"))

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect to database:", err)
	} else {
		DB.AutoMigrate(models.User{}, models.Booking{})
		log.Println("db connected")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
