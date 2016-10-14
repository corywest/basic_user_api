package main

var currentID int
var users Users

func createUser(u User) User {
	currentID++
	u.ID = currentID
	users = append(users, u)
	return u
}
