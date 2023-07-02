package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/younesabouali/rss-aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found")
	}
	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatal("dbURL is not found")
	}
	conn, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal("Couldn't connect to db")
	}
	apiCfg := apiConfig{DB: database.New(conn)}
	AppRouter(port, apiCfg)
}
