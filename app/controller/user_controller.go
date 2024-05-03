package controller

import (
	"fmt"
	"net/http"

	"go-auth/app/db"
	"go-auth/app/model"

	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController() *UserController {
	return &UserController{DB: db.DB}
}

func (uc *UserController) HandleHelloWorld(w http.ResponseWriter, r *http.Request) {
	// Create a new user
	user := model.User{Name: "John Doe", Email: "john@example.com"}
	uc.DB.Create(&user)

	// Query the database
	var retrievedUser model.User
	uc.DB.First(&retrievedUser, "name = ?", "John Doe")

	// Print response
	fmt.Fprintf(w, "Hello, %s with email %s!", retrievedUser.Name, retrievedUser.Email)
}
