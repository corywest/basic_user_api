package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		Name:        "HelloUserHandler",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: HelloUserHandler,
	},
	Route{
		Name:        "GetAllUsersHandler",
		Method:      "GET",
		Pattern:     "/users",
		HandlerFunc: GetAllUsersHandler,
	},
	Route{
		Name:        "GetUserHandler",
		Method:      "GET",
		Pattern:     "/users/{userID}",
		HandlerFunc: GetUserHandler,
	},
}
