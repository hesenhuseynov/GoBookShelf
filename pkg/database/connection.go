package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() error {
	// var err error

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Errir loading .env file")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		customErr := fmt.Errorf("Database connection failed:%v ", err)
		return customErr
	}

	return DB.Ping()
}
