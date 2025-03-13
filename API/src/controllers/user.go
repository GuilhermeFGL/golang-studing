package controllers

import (
	"database/sql"
	"encoding/json"
	"example.com/m/v2/API/src/database"
	"example.com/m/v2/API/src/models"
	"example.com/m/v2/API/src/repository"
	"example.com/m/v2/API/src/util/httpresponse"
	"io"
	"log"
	"net/http"
)

// ListUsers list all users
func ListUsers(w http.ResponseWriter, r *http.Request) {

}

// GetUser search a user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {

}

// CreateUser create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "Error reading body")
		log.Fatal("Error reading body", err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "Error parsing body")
		log.Fatal("Error parsing body", err)
		return
	}

	if err := user.Prepare(); err != nil {
		httpresponse.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	db, err := database.Connect()
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, "Error connecting to database")
		log.Fatal("Error connecting to database", err)
		return
	}

	userRepository := repository.NewRepository(db)
	createdUser, err := userRepository.CreateUser(user)
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, "Error creating user")
		log.Fatal("Error creating user", err)
		return
	}

	httpresponse.JSON(w, http.StatusCreated, createdUser)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing database connection", err)
		}
	}(db)
}

// UpdateUser update an existing user searching ny ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DeleteUser delete an existing user
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}
