package main

import (
	"net/http"
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
		"RCount",
		"GET",
		"/status/rcount",
		RCount,
	},
	Route{
		"QDepth",
		"GET",
		"/status/qdepth",
		QDepth,
	},
	// Outbox
	Route{
		"Outbox",
		"POST",
		"/outbox",
		Outbox,
	},
	// Health
	Route{
		"Health",
		"GET",
		"/health",
		Health,
	},
	Route{
		"Poison",
		"POST",
		"/poison",
		Poison,
	},
	Route{
		"Replenish",
		"POST",
		"/replenish",
		Replenish,
	},
}
