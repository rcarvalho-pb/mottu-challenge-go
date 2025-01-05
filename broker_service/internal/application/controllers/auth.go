package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/application/dtos"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var authRequest *dtos.AuthRequest
	if err = json.Unmarshal(requestBody, authRequest); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	tokenString, err := srv.Authenticate(authRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
}
