package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	DbManager "github.com/younesabouali/rss-aggregator/internal"
)

func main() {

	godotenv.Load()
	DB := DbManager.Manager()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found")
	}
	AppRouter(port, DB)
}
