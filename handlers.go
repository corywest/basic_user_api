package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there user!")
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Name: "Bobby", Email: "bob@bobby.com"},
		User{Name: "Tommy", Email: "tom@tommy.com"},
	}

	w.Header().Set("Content-Type", "application/json;charset-UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset-UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	userID := vars["userID"]
	fmt.Fprintln(w, "User page for id:", userID)
}
