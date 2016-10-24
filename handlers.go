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
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var adminUsers AdminUsers
var localUsers Users

func HelloUserHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./admin_users.json")

	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(file)
	defer file.Close()

	if err := dec.Decode(&adminUsers); err != nil {
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

	for _, adminUser := range adminUsers {
		fmt.Println(adminUser.Email)
		fmt.Println(adminUser.Password)

		if pair[0] != adminUser.Email || pair[1] != adminUser.Password {
			http.Error(w, "Not authorized. Please enter an admin user email and password", 401)
			return
		}

		if adminUser.Admin != true {
			http.Error(w, "Admin access only!", 401)
		}
	}
	fmt.Fprintf(w, "Hello there adminUser!")
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := Users{
		User{Email: "bob@bobby.com", Password: "bobby20016"},
		User{Email: "tom@tommy.com", Password: "tommy2016"},
	}

	file, err := os.Open("./admin_users.json")

	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(file)
	defer file.Close()

	if err := dec.Decode(&adminUsers); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json;charset-UTF-8")
	w.WriteHeader(http.StatusOK)

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

	for _, adminUser := range adminUsers {
		fmt.Println(adminUser.Email)
		fmt.Println(adminUser.Password)

		if pair[0] != adminUser.Email || pair[1] != adminUser.Password {
			http.Error(w, "Not authorized. Please enter an admin user email and password", 401)
			return
		}

		if adminUser.Admin != true {
			http.Error(w, "Admin access only!", 401)
		}
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset-UTF-8")
	w.WriteHeader(http.StatusOK)

	params := mux.Vars(r)

	for _, user := range localUsers {

		if paramID, err := strconv.Atoi(params["id"]); err == nil {
			if user.ID == paramID {
				json.NewEncoder(w).Encode(user)
				return
			}
		}
	}

	json.NewEncoder(w).Encode(&Users{})
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
