package main

import (
	"encoding/json"
	"net/http"
)

type Particle struct {
	Endpoint string      `json:"endpoint"`
	Return   Return      `json:"return"`
	Data     interface{} `json:"data"`
}

type Return struct {
	Success string `json:"success"`
	Fail    string `json:"fail"`
}

func Outbox(w http.ResponseWriter, r *http.Request) {
	var incomingp Particle
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&incomingp); err != nil {
		response, _ := json.Marshal("missing.particle")
		w.WriteHeader(400)
		w.Write(response)
		return
	}
	defer r.Body.Close()

	ok := false
	for !ok {
		select {
		case q <- incomingp:
			ok = true
		default:
			ok = false
		}
	}

	if ok {
		response, _ := json.Marshal("particle.received")
		w.WriteHeader(200)
		w.Write(response)
	}
}
