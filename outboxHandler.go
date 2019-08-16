package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Particle struct {
	Endpoint string `json:"endpoint"`
	RData    Return `json:"return"`
	Data     string `json:"data"`
}

type Return struct {
	Success string `json:"success"`
	Failure string `json:"failure"`
}

type q []Particle

func Outbox(w http.ResponseWriter, r *http.Request) {
	var incomingp Particle
	/*decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&incomingp); err != nil {
		response, _ := json.Marshal("missing.particle")
		w.WriteHeader(400)
		w.Write(response)
	}
	defer r.Body.Close()*/

	_ = json.NewDecoder(r.Body).Decode(&incomingp)

	fmt.Fprintln(w, "world: ", incomingp)
}

/*

curl -d '{"endpoint":"dwdw","return":{"success":"a","failure":"b"},"data":"dtt"}' -H 'Content-type: application/json' -X POST localhost:8080/outbox

*/
