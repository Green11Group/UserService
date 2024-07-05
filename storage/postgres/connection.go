package postgres

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Config() {
	err := godotenv.Load("storage/postgres/.env")
	if err != nil {
		log.Fatal("Error in file .env", err)
	}
}

func Connection() (*sql.DB, error) {
	conn := fmt.Sprintf("host%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("Postgres connection failed", err)
		return nil, err
	}

	fmt.Println(os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))

	if err := db.Ping(); err != nil {
		log.Fatal("Connection failed")
		return nil, err
	}

	return db, nil
}
