package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", router))
}
