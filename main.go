package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// create routes
	router := NewRouter()

	// start q processing
	runq()

	log.Fatal(http.ListenAndServe(":80", router))
}

func RCount(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(rCount)
}
