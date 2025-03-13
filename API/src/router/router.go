package router

import (
	"example.com/m/v2/API/src/router/route"
	"github.com/gorilla/mux"
	"log"
)

// GenerateRouter generate api router
func GenerateRouter() *mux.Router {
	log.Println("Configuring routes")

	router := mux.NewRouter()
	route.ConfigureRoutes(router)
	return router
}
