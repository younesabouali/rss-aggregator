package main

import (
	"net/http"
)

type Names struct {
	Name string `json:"name"`
}

func handleReadiness(w http.ResponseWriter, r *http.Request) {
	Myname := Names{Name: "younes"}
	respondWithJSON(w, 200, Myname)

}
func handleErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 404, "not found")
}
