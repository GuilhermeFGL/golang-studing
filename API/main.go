package main

import (
	"example.com/m/v2/API/src/router"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting API")

	r := router.GenerateRouter()

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
