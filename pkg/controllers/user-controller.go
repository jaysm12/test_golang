package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"otomo_golang/pkg/models"
	"otomo_golang/pkg/utils"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
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

type ResponseLogin struct {
	Success bool
	Token   string
}

type ResponseError struct {
	Success bool
	Msg     string
}

type ReqBodyLogin struct {
	Username string
	Password string
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

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	input := &ReqBodyLogin{}

	utils.ParseBody(r, input)

	if input.Password == "" || input.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		errorRes := ResponseError{
			Success: false,
			Msg:     "Please fill username / password",
		}

		res, _ := json.Marshal(errorRes)
		w.Write(res)
		return
	}

	user := models.FindByUsername(input.Username)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if user.User_id == 0 || err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		errorRes := ResponseError{
			Success: false,
			Msg:     "Invalid Username / Password",
		}

		res, _ := json.Marshal(errorRes)
		w.Write(res)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"user_id":  user.User_id,
	})

	tokenstring, err := token.SignedString([]byte("otomo"))

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	response := ResponseLogin{
		Success: true,
		Token:   tokenstring,
	}

	res, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id := vars["user_id"]

	user_id, err := strconv.ParseInt(id, 0, 0)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	user := models.FindByID(user_id)

	if user.User_id == 0 {
		w.WriteHeader(http.StatusNotFound)

		response := ResponseError{
			Success: false,
			Msg:     "User not found",
		}

		res, _ := json.Marshal(response)

		w.Write(res)
		return
	}

	user.DeleteUser()

	response := ResponseCreate{
		Success: true,
	}

	res, _ := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
