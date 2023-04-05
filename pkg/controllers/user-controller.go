package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"otomo_golang/pkg/models"
	"otomo_golang/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type ResponseList struct {
	Success bool
	Users   []models.User
}

type ResponseCreate struct {
	Success bool
	User    models.User
}

type ResponseError struct {
	Success bool
	Msg     string
}

type ReqBody struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := models.ListUsers()

	response := ResponseList{
		Success: true,
		Users:   users,
	}

	res, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	newUser := &models.User{}
	utils.ParseBody(r, newUser)

	if newUser.Password == "" || newUser.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorRes := ResponseError{
			Success: false,
			Msg:     "Please fill username / password",
		}

		res, _ := json.Marshal(errorRes)
		w.Write(res)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser.Password = string(hashedPassword)

	newUser.CreateUser()

	response := ResponseCreate{
		Success: true,
		User:    *newUser,
	}

	res, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
