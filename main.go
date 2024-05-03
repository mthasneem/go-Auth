package main

import (
	"fmt"
	"net/http"

	"go-auth/app/controller"
	"go-auth/app/db"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Initialize the user controller
	userController := controller.NewUserController()

	// Define a simple handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userController.HandleHelloWorld(w, r)
	})

	// Start the server on port 8080
	serverErr := http.ListenAndServe(":8080", handler)
	if serverErr != nil {
		fmt.Println("Error starting server:", serverErr)
	}
}
