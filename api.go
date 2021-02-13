package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func home(w http.ResponseWriter, r *http.Request) {
	log.Println("Called home")
	w.Write([]byte(`{"message": "To use the endpoint send REST commands against /api/v1/participant "}`))
}

func getParticipants(w http.ResponseWriter, r *http.Request) {
	listDB := make([]*Participant, 0)

	for _, participant := range db {
		fmt.Println(participant)
		listDB = append(listDB, participant)
	}

	json.NewEncoder(w).Encode(listDB)
}

func getParticipant(w http.ResponseWriter, r *http.Request) {

	id, err := getParams("id", r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	p := db[id]

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(p)
	}

}
func createParticipant(w http.ResponseWriter, r *http.Request) {
	p := Participant{}
	p.Deserialize(r)
	err := p.Validate()

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	p.ReferenceNumber = generateRandomReferenceNumber()
	db[p.ReferenceNumber] = &p
	json.NewEncoder(w).Encode(p)
}
func updateParticipant(w http.ResponseWriter, r *http.Request) {
	id, err := getParams("id", r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	p, err := searchParticipantByReferenceNumber(id)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	p.Deserialize(r)
	err = p.Validate()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	db[p.ReferenceNumber] = p
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)

}
func deleteParticipant(w http.ResponseWriter, r *http.Request) {
	id, err := getParams("id", r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	p, err := searchParticipantByReferenceNumber(id)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	delete(db, p.ReferenceNumber)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
	return
}

func generateRandomReferenceNumber() string {
	uuid := uuid.New()
	return uuid.String()
}

func searchParticipantByReferenceNumber(id string) (*Participant, error) {
	fmt.Println(id)
	fmt.Println(db)
	p := db[id]
	if p == nil {

		ErrorAPI := &ErrorAPI{
			Message:    "No Participant Found",
			HTTPStatus: http.StatusInternalServerError,
		}
		return nil, ErrorAPI
	}

	return p, nil
}
