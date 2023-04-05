package main_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"otomo_golang/pkg/controllers"
	"otomo_golang/pkg/models"
	"testing"
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

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

var createdUser *models.User
var idCreated int64

func TestCreateUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controllers.CreateUser))

	payload := []byte(`{"username":"testuser","password":"test_user", "firstName":"test", "lastName":"user"}`)

	res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(payload))

	checkResponseCode(t, http.StatusOK, res.StatusCode)

	if err != nil {
		t.Error(err)
	}

	var body ResponseCreate

	b, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(b, &body)

	createdUser := body.User

	if !body.Success {
		t.Error("Expected success should be true")
	}

	if createdUser.User_id == 0 {
		t.Error("Expected new user_id shouldn't be 0")
	}
	if createdUser.Username != "testuser" {
		t.Errorf("Expected username %s. Got %s\n", "testuser", createdUser.Username)
	}
	if createdUser.FirstName != "test" {
		t.Errorf("Expected FirstName %s. Got %s\n", "test", createdUser.FirstName)
	}
	if createdUser.Lastname != "user" {
		t.Errorf("Expected Lastname %s. Got %s\n", "user", createdUser.Lastname)
	}

}
func TestCreateUserEmptyUsername(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controllers.CreateUser))

	payload := []byte(`{"username":"","password":"test_user", "firstName":"test", "lastName":"user"}`)

	res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(payload))

	checkResponseCode(t, http.StatusBadRequest, res.StatusCode)

	if err != nil {
		t.Error(err)
	}

	var body ResponseError

	b, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(b, &body)

	if body.Success {
		t.Error("Expected Success should be false")
	}

	if body.Msg != "Please fill username / password" {
		t.Errorf("Expected error Msg %s. Got %s\n", "Please fill username / password", body.Msg)
	}

}

func TestCreateUserUsernameExist(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controllers.CreateUser))

	payload := []byte(`{"username":"testuser","password":"test_user", "firstName":"test", "lastName":"user"}`)

	res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(payload))

	checkResponseCode(t, http.StatusBadRequest, res.StatusCode)

	if err != nil {
		t.Error(err)
	}

	var body ResponseError

	b, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(b, &body)

	if body.Success {
		t.Error("Expected Success should be false")
	}

	if body.Msg != "Username already taken!" {
		t.Errorf("Expected error Msg %s. Got %s\n", "Username already taken!", body.Msg)
	}

}
func TestGetUsers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controllers.ListUsers))

	res, err := http.Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	checkResponseCode(t, http.StatusOK, res.StatusCode)

	b, _ := ioutil.ReadAll(res.Body)

	var body ResponseList

	json.Unmarshal(b, &body)

	if !body.Success {
		t.Error("Expected success should be true")
	}

	listUser := body.Users
	isExist := false

	for _, v := range listUser {
		if v.Username == "testuser" {
			isExist = true
		}
	}

	if !isExist {
		t.Error("Expected newly created user is in the list")
	}
}

func TestLogin(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controllers.Login))

	payload := []byte(`{"username":"testuser","password":"test_user"}`)

	res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(payload))
	checkResponseCode(t, http.StatusOK, res.StatusCode)

	if err != nil {
		t.Error(err)
	}

	b, _ := ioutil.ReadAll(res.Body)

	var body ResponseLogin

	json.Unmarshal(b, &body)

	if !body.Success {
		t.Error("Expected success should be true")
	}

	if body.Token == "" {
		t.Error("Expected token to exist")
	}
}

func TestLoginEmptyUsername(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controllers.Login))

	payload := []byte(`{"username":"","password":"wrongpass"}`)

	res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(payload))
	checkResponseCode(t, http.StatusBadRequest, res.StatusCode)

	if err != nil {
		t.Error(err)
	}

	var body ResponseError

	b, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(b, &body)

	if body.Success {
		t.Error("Expected Success should be false")
	}

	if body.Msg != "Please fill username / password" {
		t.Errorf("Expected error Msg %s. Got %s\n", "Please fill username / password", body.Msg)
	}
}
func TestLoginWrongPassword(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controllers.Login))

	payload := []byte(`{"username":"testuser","password":"wrongpass"}`)

	res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(payload))
	checkResponseCode(t, http.StatusUnauthorized, res.StatusCode)

	if err != nil {
		t.Error(err)
	}

	var body ResponseError

	b, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(b, &body)

	if body.Success {
		t.Error("Expected Success should be false")
	}

	if body.Msg != "Invalid Username / Password" {
		t.Errorf("Expected error Msg %s. Got %s\n", "Invalid Username / Password", body.Msg)
	}
}

func TestDeletUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controllers.DeleteUser))

	payload := []byte(`{"username":"testuser","password":"test_user"}`)

	res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(payload))
	checkResponseCode(t, http.StatusOK, res.StatusCode)

	if err != nil {
		t.Error(err)
	}

	listServer := httptest.NewServer(http.HandlerFunc(controllers.ListUsers))

	resList, _ := http.Get(listServer.URL)

	checkResponseCode(t, http.StatusOK, resList.StatusCode)

	b, _ := ioutil.ReadAll(resList.Body)

	var body ResponseList

	json.Unmarshal(b, &body)

	listUser := body.Users

	isExist := false

	for _, v := range listUser {
		if v.Username == "testuser" {
			isExist = true
		}
	}

	if isExist {
		t.Error("Expected user should be deleted")
	}
}

func TestDeleteWrongPassword(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controllers.DeleteUser))

	payload := []byte(`{"username":"testuser","password":"wrongpass"}`)

	res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(payload))
	checkResponseCode(t, http.StatusUnauthorized, res.StatusCode)

	if err != nil {
		t.Error(err)
	}

	var body ResponseError

	b, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(b, &body)

	if body.Success {
		t.Error("Expected Success should be false")
	}

	if body.Msg != "Invalid Username / Password" {
		t.Errorf("Expected error Msg %s. Got %s\n", "Invalid Username / Password", body.Msg)
	}
}

func TestDeleteEmptyUsername(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(controllers.DeleteUser))

	payload := []byte(`{"username":"","password":"wrongpass"}`)

	res, err := http.Post(server.URL, "application/json", bytes.NewBuffer(payload))
	checkResponseCode(t, http.StatusBadRequest, res.StatusCode)

	if err != nil {
		t.Error(err)
	}

	var body ResponseError

	b, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(b, &body)

	if body.Success {
		t.Error("Expected Success should be false")
	}

	if body.Msg != "Please fill username / password" {
		t.Errorf("Expected error Msg %s. Got %s\n", "Please fill username / password", body.Msg)
	}
}
