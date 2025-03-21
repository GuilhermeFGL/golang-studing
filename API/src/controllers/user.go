package controllers

import (
	"database/sql"
	"encoding/json"
	"example.com/m/v2/API/src/database"
	"example.com/m/v2/API/src/models"
	"example.com/m/v2/API/src/repository"
	"example.com/m/v2/API/src/security"
	"example.com/m/v2/API/src/util/httpresponse"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ListUsers list all users
func ListUsers(w http.ResponseWriter, r *http.Request) {
	queryName := strings.ToLower(r.URL.Query().Get("name"))

	db, err := database.Connect()
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, "Error connecting to database")
		log.Fatal("Error connecting to database", err)
		return
	}

	userRepository := repository.NewUserRepository(db)
	users, err := userRepository.SearchUser(queryName)
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, "Error creating user")
		return
	}

	if users == nil {
		httpresponse.Error(w, http.StatusNoContent, "Users not found")
	} else {
		httpresponse.JSON(w, http.StatusOK, users)
	}
}

// GetUser search users by name
func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	db, err := database.Connect()
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, "Error connecting to database")
		log.Fatal("Error connecting to database", err)
		return
	}

	userRepository := repository.NewUserRepository(db)
	user, err := userRepository.FetchUser(userID)
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, "Error fetching user: "+err.Error())
		return
	}

	if user.ID == 0 {
		httpresponse.Error(w, http.StatusNotFound, "User not found")
	} else {
		httpresponse.JSON(w, http.StatusOK, user)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing db", err)
		}
	}(db)
}

// CreateUser create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "Error reading body")
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "Error parsing body")
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

	userRepository := repository.NewUserRepository(db)
	createdUser, err := userRepository.CreateUser(user)
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, "Error creating user")
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
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	tokenUserID, err := security.ExtractUserIdFromToken(r)
	if err != nil || tokenUserID == 0 || tokenUserID != userID {
		httpresponse.Error(w, http.StatusUnauthorized, "User not authorized to update user with id "+strconv.FormatUint(userID, 10))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "Error reading body")
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "Error parsing body")
		return
	}

	if err := user.PrepareForUpdate(); err != nil {
		httpresponse.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	db, err := database.Connect()
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, "Error connecting to database")
		log.Fatal("Error connecting to database", err)
		return
	}

	userRepository := repository.NewUserRepository(db)
	updatedUser, err := userRepository.UpdateUser(userID, user)
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, "Error updating user")
		return
	}

	if updatedUser.ID == 0 {
		httpresponse.Error(w, http.StatusNotFound, "User not found")
	} else {
		httpresponse.JSON(w, http.StatusOK, updatedUser)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing database connection", err)
		}
	}(db)
}

// DeleteUser delete an existing user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	tokenUserID, err := security.ExtractUserIdFromToken(r)
	if err != nil || tokenUserID == 0 || tokenUserID != userID {
		httpresponse.Error(w, http.StatusUnauthorized, "User not authorized to update user with id "+strconv.FormatUint(userID, 10))
		return
	}

	db, err := database.Connect()
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, "Error connecting to database")
		log.Fatal("Error connecting to database", err)
		return
	}

	userRepository := repository.NewUserRepository(db)
	deletedUser, err := userRepository.DeleteUser(userID)
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, "Error deleting user")
		return
	}

	if deletedUser {
		httpresponse.Error(w, http.StatusOK, "User deleted")
	} else {
		httpresponse.Error(w, http.StatusNotFound, "User not found")
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing database connection", err)
		}
	}(db)
}

// FollowUser allow user to follow another user
func FollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	followerID, err := strconv.ParseUint(parameters["followerId"], 10, 64)
	if err != nil {
		httpresponse.Error(w, http.StatusBadRequest, "Invalid follower ID")
		return
	}

	tokenUserID, err := security.ExtractUserIdFromToken(r)
	if err != nil || tokenUserID == 0 || tokenUserID != userID {
		httpresponse.Error(w, http.StatusUnauthorized, "User not authorized to update user with id "+strconv.FormatUint(userID, 10))
		return
	}

	if userID == followerID {
		httpresponse.Error(w, http.StatusBadRequest, "Users can't follow themself")
		return
	}

	db, err := database.Connect()
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, "Error connecting to database")
		log.Fatal("Error connecting to database", err)
		return
	}

	userRepository := repository.NewUserRepository(db)

	user, err := userRepository.FetchUser(userID)
	if err != nil || user.ID == 0 {
		httpresponse.Error(w, http.StatusNotFound, "User not found")
		return
	}

	follower, err := userRepository.FetchUser(followerID)
	if err != nil || follower.ID == 0 {
		httpresponse.Error(w, http.StatusNotFound, "Follower not found")
		return
	}

	if err = userRepository.Follow(userID, followerID); err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, "Error following user")
		return
	}

	httpresponse.JSON(w, http.StatusNoContent, nil)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing database connection", err)
		}
	}(db)
}
