package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func readRequestBody(r *http.Request) interface{} {
	// Read a body for the request if there is one, else retur nan empty interface{}
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

func interfaceToMap(b interface{}) (map[string]interface{}, bool) {
	m, ok := b.(map[string]interface{})
	if !ok {
		ErrorLogger.Println("Could not read the JSON body.")
		return nil, true
	}
	for k, v := range m {
		fmt.Println(k, "=>", v)
	}
	return m, false
}
