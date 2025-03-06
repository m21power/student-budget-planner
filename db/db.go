package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func ConnectDB() (*sql.DB, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file: %v", err)
	// 	return nil, err
	// }
	if _, err := os.Stat(".env"); err == nil {
		if loadErr := godotenv.Load(); loadErr != nil {
			log.Printf("Warning: Could not load .env file: %v", loadErr)
		}
	}
	var user = os.Getenv("DB_USER")
	var password = os.Getenv("DB_PASSWORD")
	var dbname = os.Getenv("DB_NAME")
	var host = os.Getenv("DB_HOST")
	var port = os.Getenv("DB_PORT")
	var sslmode = os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Could not open database: %v", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return nil, err
	}

	log.Println("Connected to PostgreSQL successfully")
	return db, nil
}
