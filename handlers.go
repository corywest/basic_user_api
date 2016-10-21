package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func HelloUserHandler(w http.ResponseWriter, r *http.Request) {
	var users Users

	file, err := os.Open("./users.json")

	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(file)
	defer file.Close()

	if err := dec.Decode(&users); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		http.Error(w, "Not authorized", 401)
		return
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		http.Error(w, "Not authorized", 401)
		return
	}

	for _, user := range users {
		fmt.Println("Inside for loop...")
		fmt.Println(user.Email)
		fmt.Println(user.Password)

		if pair[0] != user.Email || pair[1] != user.Password {
			http.Error(w, "Not authorized. Please enter a user email and password", 401)
			return
		}

		if user.Admin != true {
			http.Error(w, "Admin access only!", 403)
		}
	}
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
