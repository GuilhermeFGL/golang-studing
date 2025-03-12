package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
)

type user struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading body"))
		return
	}

	var newUser user
	err = json.Unmarshal(requestBody, &newUser)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error unmarshalling body"))
	}

	db, err := Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	statement, err := db.Prepare("INSERT INTO user(name, email) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error preparing statement"))
		return
	}

	inserted, err := statement.Exec(newUser.Name, newUser.Email)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error executing statement"))
		return
	}

	id, err := inserted.LastInsertId()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting last inserted id"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user{Id: uint32(id), Name: newUser.Name, Email: newUser.Email})

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(statement)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	query, err := db.Query("SELECT id, name, email FROM user")
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error recovering users"))
		return
	}

	var users []user
	for query.Next() {
		var user user
		err = query.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error scanning user"))
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(query)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error parsing id"))
		return
	}

	db, err := Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	query, err := db.Query("SELECT id, name, email FROM user WHERE id=?", ID)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error preparing query"))
		return
	}

	var foundUser user
	if query.Next() {
		if err := query.Scan(&foundUser.Id, &foundUser.Name, &foundUser.Email); err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error parsing query"))
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundUser)

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(query)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error parsing id"))
		return
	}

	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading body"))
		return
	}

	var newUser user
	err = json.Unmarshal(requestBody, &newUser)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error unmarshalling body"))
	}

	db, err := Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	query, err := db.Query("SELECT id, name, email FROM user WHERE id=?", ID)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error preparing query"))
		return
	}

	if !query.Next() {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	statement, err := db.Prepare("UPDATE user SET name=?, email=? WHERE id=?")
	updated, err := statement.Exec(newUser.Name, newUser.Email, ID)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error executing statement"))
		return
	}

	_, err = updated.LastInsertId()
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error getting last updated id"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user{Id: uint32(ID), Name: newUser.Name, Email: newUser.Email})

	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(query)

	defer func(statement *sql.Stmt) {
		err := statement.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(statement)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error parsing id"))
		return
	}

	db, err := Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	deleted, err := db.Exec("DELETE FROM user WHERE id=?", ID)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error executing delete"))
		return
	}

	if rows, _ := deleted.RowsAffected(); rows > 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
}
