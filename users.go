package main

import (
	"fmt"
	"net/http"
)

func handleUsers(w http.ResponseWriter, req *http.Request) {
	// Receiving the body to pass along
	body := readRequestBody(req)
	// Handle the types of requests
	switch req.Method {
	case "GET":
		usersGET(body)
	case "POST":
		usersPOST()
	case "PUT":
		usersPUT()
	case "DELETE":
		usersDELETE()
	}
}

func usersGET(body interface{}) {
	m, ok := body.(map[string]interface{})
	if !ok {
		ErrorLogger.Println("Could not read the JSON body.")
		fmt.Println("OOPS")
	}
	for k, v := range m {
		fmt.Println(k, "=>", v)
	}
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
