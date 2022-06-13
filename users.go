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
	m, test := interfaceToMap(body)
	if test {
		ErrorLogger.Println("Failed to convert JSON body")
	}
	str := "SELECT "
	if fields, ok := m["fields"]; ok {
		if s, ok := fields.([]interface{}); ok {
			for idx, val := range s {
				str += fmt.Sprintf("%s", val)
				if (idx + 1) != len(s) {
					str += ", "
				}
			}
			str += " "
		} else {
			ErrorLogger.Println("Failed to use fields slice")
		}
	} else {
		str += "* "
	}
	str += "FROM users"
	if fields, ok := m["where"]; ok {
		if s, ok := fields.(map[string]interface{}); ok {
			fmt.Println(s)
			for idx, val := range s {
				fmt.Println(idx, val)
			}
		} else {
		}
	} else {
	}
	fmt.Println(str)
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
