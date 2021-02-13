package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Participant : Object Definition for Participants
type Participant struct {
	Name            string `json:"name" validate:"required"`
	ReferenceNumber string `json:"referenceNumber" `
	DateOfBirth     string `json:"dateOfBirth" validate:"required"`
	PhoneNumber     string `json:"phoneNumber" validate:"required"`
	Address         string `json:"address" validate:"required"`
}

func (p *Participant) Deserialize(r *http.Request) (err error) {

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&p)

	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

func (p *Participant) Validate() error {

	v := validator.New()
	valErr := v.Struct(p)

	if valErr != nil {
		ErrorAPI := &ErrorAPI{
			Message:    valErr.Error(),
			HTTPStatus: http.StatusNotAcceptable,
		}
		return ErrorAPI
	}
	return nil

}
