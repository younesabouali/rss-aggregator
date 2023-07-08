package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	DbManager "github.com/younesabouali/rss-aggregator/internal"
	"github.com/younesabouali/rss-aggregator/scrapper"
)

func main() {

	godotenv.Load()
	DB := DbManager.Manager()
	feedScrap := scrapper.FeedScrapper{DB: DB}
	go feedScrap.Scrape(5, 1*time.Minute)
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found")
	}
	AppRouter(port, DB)
}
