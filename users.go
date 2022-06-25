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
		// Write a JSON response
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(response)
		errorHandle(err, "Failed to convert JSON body")
		w.Write(j)
	case "POST":
		// add a user
		usersPOST(fmt.Sprintf("%v", req.URL))
	case "PUT":
		usersPUT()
	case "DELETE":
		usersDELETE()
	}
}

func usersGET(body interface{}) []interface{} {

	// Ensure a map[string]interface{} to handle the retrieval of users.
	m, test := interfaceToMap(body)
	if test {
		ErrorLogger.Println("Failed to convert JSON body")
	}
	// Prepare the query string.
	var stmnt string
	if id, ok := m["id"]; ok {
		// Single ID GET.
		stmnt = fmt.Sprintf("SELECT * FROM users WHERE id = %d;", int(id.(float64)))
	} else if ids, ok := m["ids"]; ok {
		// Multiple IDs query.
		// Format the slice, into a joined string that can be used in the query.
		var s []string
		for _, v := range ids.([]interface{}) {
			s = append(s, fmt.Sprintf("%d", int(v.(float64))))
		}
		arrStr := strings.Join(s, ",")
		stmnt = fmt.Sprintf("SELECT * FROM users WHERE id IN (%s);", arrStr)
	} else {
		// Base select all users.
		stmnt = "SELECT * FROM users;"
	}
	// Query the DB to fetch the data.
	rows, err := DB.Query(stmnt)
	errorHandle(err, "Failed to query DB")
	var resp []interface{}
	// Loop over all the rows.
	for rows.Next() {
		var u User
		err = rows.Scan(&u.Id, &u.Name, &u.Surname, &u.Age)
		errorHandle(err, "Failed to read the row values")
		// Append the User Struct to the response array.
		resp = append(resp, structToMap(u))
	}
	return resp
}

func usersPOST(u string) (map[string]interface{}, string) {
	// Get the action specified from the URL
	split := strings.Split(u, "/")
	action := split[len(split)-1]
	// Should be add only
	if action != "add" {
		return nil, "No such supported action"
	}
	fmt.Println("POST USERS", action)
	return nil, ""
}

func usersPUT() {
	fmt.Println("PUT USERS")
}

func usersDELETE() {
	fmt.Println("DELETE USERS")
}
