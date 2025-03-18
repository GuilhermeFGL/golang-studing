package route

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Uri           string
	Method        string
	Func          func(http.ResponseWriter, *http.Request)
	AuthRequested bool
}

// ConfigureRoutes configure and add routes
func ConfigureRoutes(router *mux.Router) {
	routes := UserRoutes
	routes = append(routes, LoginRoutes...)

	for _, route := range routes {
		router.HandleFunc(route.Uri, route.Func).Methods(route.Method)
	}
}
