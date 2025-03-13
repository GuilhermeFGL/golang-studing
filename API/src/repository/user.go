package repository

import (
	"database/sql"
	"example.com/m/v2/API/src/models"
	"log"
)

type User struct {
	db *sql.DB
}

// NewRepository create a new repository for user
func NewRepository(db *sql.DB) *User {
	return &User{db}
}

// CreateUser create and return a new user
func (repository User) CreateUser(user models.User) (models.User, error) {
	statement, err := repository.db.Prepare("INSERT INTO users (name, email, nickname, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return models.User{}, err
	}

	exec, err := statement.Exec(user.Name, user.Email, user.NickName, user.Password)
	if err != nil {
		return models.User{}, err
	}

	user.ID, err = exec.LastInsertId()

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Fatal("Error closing statement")
		}
	}(statement)

	return user, nil
}
