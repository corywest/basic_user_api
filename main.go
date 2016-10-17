package main

import (
	"log"
	"net/http"

	"github.com/goji/httpauth"
)

func main() {
	router := NewRouter()

	http.Handle("/", httpauth.SimpleBasicAuth("dave", "somepassword")(router))
	log.Fatal(http.ListenAndServe(":8080", router))
}
