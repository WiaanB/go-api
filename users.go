package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func (u *User) validateUser() []string {
	var s []string
	if u.Name == "" {
		s = append(s, "no name supplied")
	}
	if u.Age < 0 {
		s = append(s, "user needs an age above 0")
	}
	if u.Surname == "" {
		s = append(s, "no surname supplied")
	}
	return s
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
		w.Header().Set("Content-Type", "application/json")
		// add a user
		val, msg := usersPOST(fmt.Sprintf("%v", req.URL), body)
		if msg != "" {
			j, err := json.Marshal(map[string]interface{}{"error": msg})
			errorHandle(err, "Failed to convert JSON body")
			w.Write(j)
		} else {
			j, err := json.Marshal(val)
			errorHandle(err, "Failed to convert JSON body")
			w.Write(j)
		}
	case "PUT":
		usersPUT(fmt.Sprintf("%v", req.URL), body)
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

func usersPOST(u string, body interface{}) (map[string]interface{}, string) {
	// Get the action specified from the URL
	split := strings.Split(u, "/")
	action := split[len(split)-1]
	// Should be add only
	if action != "add" {
		return nil, "unsupported action"
	}
	// Create a Struct from the body
	m, test := interfaceToMap(body)
	if test {
		ErrorLogger.Println("Failed to convert JSON body")
	}
	jsonStr, err := json.Marshal(m)
	errorHandle(err, "Failed to marshal JSON")
	// create User
	var user User
	err = json.Unmarshal(jsonStr, &user)
	errorHandle(err, "Failed to marshal JSON")
	errors := user.validateUser()
	if len(errors) > 0 {
		return map[string]interface{}{"errors": errors}, ""
	} else {
		err = DB.QueryRow("INSERT INTO users (name, surname, age) VALUES($1, $2, $3) RETURNING id", user.Name, user.Surname, user.Age).Scan(&user.Id)
		errorHandle(err, "Failed to add user")
	}
	return map[string]interface{}{"status": 200, "message": "user created successfully", "data": user}, ""
}

func usersPUT(url string, body interface{}) (map[string]interface{}, string) {
	// Get the id specified from the URL
	split := strings.Split(url, "/")
	id := split[len(split)-1]
	// cast the id to string
	givenInt, err := strconv.Atoi(id)
	if err != nil {
		errorHandle(err, "Failed to convert id to int")
		return nil, "failed to convert id to int"
	}
	// Create a Struct from the body
	m, test := interfaceToMap(body)
	if test {
		ErrorLogger.Println("Failed to convert JSON body")
	}
	DB.QueryRow("UPDATE users SET name = $2, surname = $3, age = $4 WHERE id = $1", givenInt, m["name"], m["surname"], m["age"])
	fmt.Println(givenInt, m)
	return nil, ""
}

func usersDELETE() {
	fmt.Println("DELETE USERS")
}
