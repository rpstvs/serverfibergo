package sql

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rpstvs/serverfibergo/database"
)

func CreateDBInstance() *database.Queries {
	godotenv.Load(".env")
	dbURL := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Printf("Couldnt connect to database")
	}
	DB := database.New(db)

	return DB
}
