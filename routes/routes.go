package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route ...
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

// Routes ...
type Routes []Route

// NewRouter ...
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandleFunc)
	}
	return router
}

var routes = Routes{
	Route{"Index", "POST", "/reservations/{reservationId}/email-confirmation", EmailConfirmation},
}
