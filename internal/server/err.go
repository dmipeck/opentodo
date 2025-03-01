package server

import (
	"log"
	"net/http"
)

func requestErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func responseErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("error: %v", err)
	http.Error(w, "internal server error", http.StatusInternalServerError)
}
