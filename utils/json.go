package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func BodyParser[T interface{}](r *http.Request, e T) (T, error) {
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		return e, errors.New("Couldn't parse params")
	}
	return e, nil

}
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responded with 500 error ", msg)
	}
	type responseError struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, responseError{Error: msg})

}
