package controllers

import (
	"encoding/json"
	"net/http"
	"otomo_golang/pkg/models"
)

type ResponseUser struct {
	Success bool
	Users   []models.User
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := models.ListUsers()

	response := ResponseUser{
		Success: true,
		Users:   users,
	}

	res, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
