package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Read a body for the request if there is one, else return an empty interface{}
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

// Try to cast a interface into a JSON map[string]interface{}
func interfaceToMap(b interface{}) (map[string]interface{}, bool) {
	if m, ok := b.(map[string]interface{}); !ok {
		ErrorLogger.Println("Could not read the JSON body.")
		return nil, true
	} else {
		return m, false
	}
}
