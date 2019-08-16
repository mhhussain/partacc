package main

import (
	"encoding/json"
	"net/http"
)

type HealthObj struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Health(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h)
}

func Poison(w http.ResponseWriter, r *http.Request) {
	h.Status = "poisoned"
	h.Message = "poisoned"
	json.NewEncoder(w).Encode(h)
}

func Replenish(w http.ResponseWriter, r *http.Request) {
	h.Status = "listening"
	h.Message = "replenished"
	json.NewEncoder(w).Encode(h)
}

var h = HealthObj{
	"partacc",
	"listening",
	"listening",
}
