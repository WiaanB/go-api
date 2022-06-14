package main

import (
	"fmt"
	"net/http"
	"strings"
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

	if id, ok := m["id"]; ok {
		stmnt := fmt.Sprintf("SELECT * FROM users WHERE id = %d;", int(id.(float64)))
		fmt.Println(stmnt)
	} else if ids, ok := m["ids"]; ok {
		var s []string
		for _, v := range ids.([]interface{}) {
			s = append(s, fmt.Sprintf("%d", int(v.(float64))))
		}
		arrStr := strings.Join(s, ",")
		stmnt := fmt.Sprintf("SELECT * FROM users WHERE id IN (%s);", arrStr)
		fmt.Println(stmnt)
	} else {
		stmnt := "SELECT * FROM users;"
		fmt.Println(stmnt)
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
