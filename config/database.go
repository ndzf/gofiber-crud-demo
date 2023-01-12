package config

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	p := GetEnv("DB_PORT")
	port, err := strconv.Atoi(p)
	if err != nil {
		message := fmt.Sprintf("ERROR <convert DB_PORT>: %v", err.Error())
		log.Fatal(message)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", GetEnv("DB_HOST"), port, GetEnv("DB_USER"), GetEnv("DB_PASS"), GetEnv("DB_NAME"))

	Database, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}
	DB = Database

	fmt.Printf("Connecting to %s", Database.Name())
}
