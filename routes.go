package main

import (
	"net/http"
)

type Method string

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		GET,
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		GET,
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		GET,
		"/todos/{todoId}",
		TodoShow,
	},
	Route{
		"TodoDelete",
		DELETE,
		"/todos/{todoId}",
		TodoDestroy,
	},
	Route{
		"TodoCreate",
		POST,
		"/todos",
		TodoCreate,
	},
}
