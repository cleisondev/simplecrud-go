package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
//Get .env variables and initliaze db connection
func SetupDB() *sql.DB {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	dbConnection, err := sql.Open("postgres", connString)

	if err != nil {
		log.Fatalf("Erro ao abrir conexão com o banco: %s", err)
	}

	err = dbConnection.Ping()

	if err != nil {
		log.Fatalf("Erro ao conectar com o banco: %s", err)
	}

	fmt.Println("Successfully connected to database")

	return dbConnection
}