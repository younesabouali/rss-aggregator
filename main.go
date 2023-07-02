package main

import (
	"fmt"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	// "github.com/joho/godotenv"
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
	router := chi.NewRouter()
	router.Use(
		cors.Handler(
			cors.Options{
				AllowedOrigins:   []string{"https://*", "http://*"},
				AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowedHeaders:   []string{"*"},
				ExposedHeaders:   []string{"Link"},
				AllowCredentials: false,
				MaxAge:           300,
			},
		),
	)
	v1Router := chi.NewRouter()
	v1Router.Get("/health", handleReadiness)
	v1Router.Get("/err", handleErr)
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		log.Fatal("dbURL is not found")
	}
	conn, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatal("Couldn't connect to db")
	}
	fmt.Println("server runing on PORT : ", port)
	srv.ListenAndServe()
	apiCfg := apiConfig{DB: database.New(conn)}
	AppRouter(port, apiCfg)
}
