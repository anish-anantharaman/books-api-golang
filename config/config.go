package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)




func InitDB() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("MYSQL_DB_USER"),
		os.Getenv("MYSQL_DB_PASSWORD"),
		os.Getenv("MYSQL_DB_HOST"),
		os.Getenv("MYSQL_DB_PORT"),
		os.Getenv("MYSQL_DB_NAME"),
	)

	//open connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database : %v", err)
	}

	//verify the connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed : %v", err)
	}

	log.Println("Database connection established successfully !")
	return db
}