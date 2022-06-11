package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

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
