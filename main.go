package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var db = make(map[string]*Participant)

func main() {
	r := mux.NewRouter()
	r.Use(middleware)

	r.HandleFunc("/", home)
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/participant", getParticipants).Methods(http.MethodGet)
	api.HandleFunc("/participant/", getParticipants).Methods(http.MethodGet)
	api.HandleFunc("/participant/{id}", getParticipant).Methods(http.MethodGet)
	api.HandleFunc("/participant/", createParticipant).Methods(http.MethodPost)
	api.HandleFunc("/participant/{id}", updateParticipant).Methods(http.MethodPut)
	api.HandleFunc("/participant/{id}", deleteParticipant).Methods(http.MethodDelete)

	server := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	http.Handle("/", r)
	log.Println("Starting REST Server on :8080")
	log.Fatal(server.ListenAndServe())
}
