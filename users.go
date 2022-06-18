package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	Id      int
	Name    string
	Surname string
	Age     int
}

func handleUsers(w http.ResponseWriter, req *http.Request) {
	// Receiving the body to pass along
	body := readRequestBody(req)
	// Handle the types of requests
	switch req.Method {
	case "GET":
		response := usersGET(body)
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(response)
		if err != nil {
			ErrorLogger.Println("Failed to convert JSON body")
		}
		w.Write(j)
	case "POST":
		usersPOST()
	case "PUT":
		usersPUT()
	case "DELETE":
		usersDELETE()
	}
}

func usersGET(body interface{}) []interface{} {
	m, test := interfaceToMap(body)
	if test {
		ErrorLogger.Println("Failed to convert JSON body")
	}

	var stmnt string
	if id, ok := m["id"]; ok {
		stmnt = fmt.Sprintf("SELECT * FROM users WHERE id = %d;", int(id.(float64)))
		fmt.Println(stmnt)
	} else if ids, ok := m["ids"]; ok {
		var s []string
		for _, v := range ids.([]interface{}) {
			s = append(s, fmt.Sprintf("%d", int(v.(float64))))
		}
		arrStr := strings.Join(s, ",")
		stmnt = fmt.Sprintf("SELECT * FROM users WHERE id IN (%s);", arrStr)
		fmt.Println(stmnt)
	} else {
		stmnt = "SELECT * FROM users;"
		fmt.Println(stmnt)
	}
	rows, err := DB.Query(stmnt)
	var resp []interface{}
	for rows.Next() {
		var r User
		err = rows.Scan(&r.Id, &r.Name, &r.Surname, &r.Age)
		if err != nil {
			ErrorLogger.Println("Failed to read DB response")
		}
		resp = append(resp, structToMap(r))
	}
	return resp
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
