package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Read a body for the request if there is one, else return an empty interface{}.

func readRequestBody(r *http.Request) interface{} {
	var body interface{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorLogger.Println("Failed to verify JSON body")
		log.Fatal(err)
	} else {
		if len(b) > 0 {
			err = json.Unmarshal(b, &body)
			if err != nil {
				ErrorLogger.Println("Failed to verify JSON body")
			}
		}
	}
	return body
}

// Try to cast a interface into a JSON map[string]interface{}.

func interfaceToMap(b interface{}) (map[string]interface{}, bool) {
	if m, ok := b.(map[string]interface{}); !ok {
		ErrorLogger.Println("Could not read the JSON body.")
		return nil, true
	} else {
		return m, false
	}
}

// Takes a interface (struct) and returns the map value of it.

func structToMap(i interface{}) (m map[string]interface{}) {
	data, _ := json.Marshal(i)
	err := json.Unmarshal(data, &m)
	if err != nil {
		ErrorLogger.Println("Failed to convert JSON")
	}
	return
}

// Function to check errors and handle them

func errorHandle(e error, m string) {
	if e != nil {
		if e.Error() != "" {
			ErrorLogger.Printf("%s: %s\n", m, e.Error())
			log.Fatalf("%s: %s\n", m, e.Error())
		} else {
			ErrorLogger.Println(m)
			log.Fatalf("%s\n", m)
		}
	}
}
