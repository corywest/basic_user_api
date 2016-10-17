package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there user!")
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Email: "bob@bobby.com", Password: "bobby20016"},
		User{Email: "tom@tommy.com", Password: "tommy2016"},
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

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json;charset-UTF-8")
		w.WriteHeader(422)
	}

	u := createUser(user)
	w.Header().Set("Content-Type", "application/json;charset-UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(u); err != nil {
		panic(err)
	}
}
