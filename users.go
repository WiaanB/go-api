package main

import (
	"fmt"
	"net/http"
)

func handleUsers(w http.ResponseWriter, req *http.Request) {
	// Printing the Request
	switch req.Method {
	case "GET":
		usersGET()
	case "POST":
		usersPOST()
	case "PUT":
		usersPUT()
	case "DELETE":
		usersDELETE()
	}
}

func usersGET() {
	fmt.Println("GET USERS")
}

func usersPOST() {
	fmt.Println("POST USERS")
}

func usersPUT() {
	fmt.Println("PUT USERS")
}

func usersDELETE() {
	fmt.Println("DELETE USERS")
}
