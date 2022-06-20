package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
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
		usersPOST(fmt.Sprintf("%v", req.URL))
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
	} else if ids, ok := m["ids"]; ok {
		var s []string
		for _, v := range ids.([]interface{}) {
			s = append(s, fmt.Sprintf("%d", int(v.(float64))))
		}
		arrStr := strings.Join(s, ",")
		stmnt = fmt.Sprintf("SELECT * FROM users WHERE id IN (%s);", arrStr)
	} else {
		stmnt = "SELECT * FROM users;"
	}
	rows, err := DB.Query(stmnt)
	if err != nil {
		ErrorLogger.Println("Failed to query DB")
	}
	var resp []interface{}
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.Surname, &u.Age)
		if err != nil {
			ErrorLogger.Println("Failed to read DB response")
		}
		resp = append(resp, structToMap(u))
	}
	return resp
}

func usersPOST(u string) {
	split := strings.Split(u, "/")
	action := split[len(split)-1]
	if action == "" {
		fmt.Println("EMPTY?!")
	}
	fmt.Println("POST USERS", action)
}

func usersPUT() {
	fmt.Println("PUT USERS")
}

func usersDELETE() {
	fmt.Println("DELETE USERS")
}
