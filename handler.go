package main

import (
	"net/http"
)

type Names struct {
	Name string `json:"name"`
}
type MyController struct {
	Collection string
}

func (e MyController) search(w http.ResponseWriter, r *http.Request) {}
func handleReadiness(w http.ResponseWriter, r *http.Request) {
	myc := MyController{}
	myc.search(w, r)
	Myname := Names{Name: "younes"}
	respondWithJSON(w, 200, Myname)

}
func handleErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 404, "not found")
}
