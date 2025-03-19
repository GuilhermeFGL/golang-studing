package controllers

import (
	"database/sql"
	"encoding/json"
	"example.com/m/v2/API/src/database"
	"example.com/m/v2/API/src/models"
	"example.com/m/v2/API/src/repository"
	"example.com/m/v2/API/src/security"
	"example.com/m/v2/API/src/util/httpresponse"
	"io"
	"log"
	"net/http"
)

// Login authenticate user
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		httpresponse.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		httpresponse.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	db, err := database.Connect()
	userRepository := repository.NewUserRepository(db)
	foundUser, err := userRepository.FindByEmail(user.Email)
	if err != nil {
		httpresponse.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	} else if foundUser.ID == 0 {
		httpresponse.Error(w, http.StatusNotFound, "User not found")
		return
	}

	if !security.CheckPasswordHash(foundUser.Password, user.Password) {
		httpresponse.Error(w, http.StatusUnauthorized, "Incorrect password")
		return
	}

	token, err := security.CreateToken(foundUser.ID)
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, err.Error())
	}
	httpresponse.JSON(w, http.StatusOK, models.Token{
		Token: token,
	})

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing database connection")
		}
	}(db)
}
