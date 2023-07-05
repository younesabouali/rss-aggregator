package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/younesabouali/rss-aggregator/Controllers"
	DbManager "github.com/younesabouali/rss-aggregator/internal"
)

func AppRouter(port string) {

	DB := DbManager.Manager()
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
	// userController := Controllers.UserController{DB: DB}
	// v1Router.Get("/users", userController.Seed)
	v1Router.Mount("/feed_follows", Controllers.FollowRouter(DB))
	v1Router.Mount("/users", Controllers.UserRouter(DB))
	v1Router.Mount("/feeds", Controllers.FeedRouter(DB))
	router.Mount("/v1", v1Router)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	fmt.Println("server runing on PORT : ", port)
	srv.ListenAndServe()
}
