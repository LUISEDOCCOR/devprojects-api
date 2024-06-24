package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conn() {

	_ = godotenv.Load()

	HOST := os.Getenv("DB_HOST")
	USER := os.Getenv("DB_USER")
	PASSWORD := os.Getenv("DB_PASSWORD")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_DBNAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		HOST, USER, PASSWORD, DBNAME, PORT)

	var error error
	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal(error.Error())
	} else {
		fmt.Println("The database is 🔥")
	}
}
