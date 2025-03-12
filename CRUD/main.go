package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // implicit import
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user", CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user", ListUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", GetUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
