package models

import (
	"otomo_golang/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastName"`
}

func init() {
	config.Connect()

	db = config.GetDB()

	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)

	return u
}

func ListUsers() []User {
	var Users []User

	db.Find(&Users)

	return Users
}

func FindByID(ID int64) *User {

	var user User
	db.Where("user_id = ?", ID).Find(&user)

	return &user
}

func (u *User) DeleteUser() {
	db.Delete(&u)
}
