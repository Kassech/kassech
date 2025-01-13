package database

import (
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var REDIS *redis.Client
var DB *gorm.DB

// Connect initializes the database connection and stores it in the global DB variable.
func Connect() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbSslMode := os.Getenv("DB_SSLMODE")
	dbTimeZone := os.Getenv("DB_TIMEZONE")

	// Ensure all required environment variables are provided
	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" || dbPort == "" || dbSslMode == "" || dbTimeZone == "" {
		log.Fatal("one or more required environment variables are missing")
	}

	// Data Source Name (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSslMode, dbTimeZone)

	// Open a connection to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}
	// Optionally, you can log the successful connection
	log.Println("Successfully connected to the database")
	REDIS = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	log.Println("Successfully connected to the redis cache")

}
