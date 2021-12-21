package router

import (
	"net/http"

	handler "example.com/logos106/saroop-api/handlers"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"GetAdmins",
		"GET",
		"/admins",
		handler.GetAdmins,
	},
	Route{
		"GetAdmin",
		"GET",
		"/admin/{id}",
		handler.GetAdmin,
	},
}
