package main

import (
	"net/http"

	utils "github.com/younesabouali/rss-aggregator/utils"
)

type Names struct {
	Name string `json:"name"`
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	Myname := Names{Name: "younes"}
	utils.RespondWithJSON(w, 200, Myname)

}
func handleErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 404, "not found")
}
