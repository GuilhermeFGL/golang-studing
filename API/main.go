package main

import (
	"example.com/m/v2/API/src/configuration"
	"example.com/m/v2/API/src/router"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting API")

	configuration.LoadConfigurations()

	r := router.GenerateRouter()

	log.Printf("Listening on port %s /n", configuration.ApiPort)
	log.Fatal(http.ListenAndServe(":"+configuration.ApiPort, r))
}
