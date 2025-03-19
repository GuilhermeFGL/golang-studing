package route

import (
	"example.com/m/v2/API/src/security"
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

	for _, r := range routes {
		if r.AuthRequested {
			router.HandleFunc(r.Uri, security.Logger(security.Authorize(r.Func))).Methods(r.Method)
		} else {
			router.HandleFunc(r.Uri, security.Logger(r.Func)).Methods(r.Method)
		}

		router.HandleFunc(r.Uri, r.Func).Methods(r.Method)
	}
}
