package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if (r.Method == "POST" || r.Method == "PUT") && r.Header.Get("Content-type") != "application/json" {
			ErrorAPI := &ErrorAPI{
				Message:    "Unsupported Media Type. Only JSON files are allowed",
				HTTPStatus: http.StatusUnsupportedMediaType,
			}
			w.WriteHeader(ErrorAPI.HTTPStatus)
			json.NewEncoder(w).Encode(ErrorAPI)
			return
		}

		if !(r.Method == "POST" || r.Method == "GET" || r.Method == "PUT" || r.Method == "DELETE") {
			ErrorAPI := &ErrorAPI{
				Message:    "Method Not Allowed: Only GET/POST/PUT/DELETE",
				HTTPStatus: http.StatusMethodNotAllowed,
			}
			w.WriteHeader(ErrorAPI.HTTPStatus)
			json.NewEncoder(w).Encode(ErrorAPI)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// Pass control back to the handler
		handler.ServeHTTP(w, r)
	})
}
func getParams(p string, r *http.Request) (string, error) {

	pathParam := mux.Vars(r)

	param := pathParam[p]
	if param == "" {
		errorMessage := fmt.Sprintf("Parameter [%s] Not Found", p)
		ErrorAPI := &ErrorAPI{
			Message:    errorMessage,
			HTTPStatus: http.StatusInternalServerError,
		}
		return "", ErrorAPI
	}

	return param, nil
}
