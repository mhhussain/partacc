package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// create routes
	router := NewRouter()

	// start q processing
	runq()

	log.Fatal(http.ListenAndServe(":91", router))
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
