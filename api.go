package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// 5mb log file size limit, adjust to your liking
const FILE_SIZE_LIMIT = 5242880

var (
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	InfoLogger    *log.Logger
)

func init() {
	fi, err := os.Stat("logs.txt")
	if err != nil {
		// handle any errors besides the log file maybe not existing
		if !strings.Contains(err.Error(), "no such file or directory") {
			log.Fatal(err)
		}
	}

	// if the file exists, and it's too large, delete it
	if fi != nil && fi.Size() > FILE_SIZE_LIMIT {
		err = os.Remove("logs.txt")
		if err != nil {
			log.Fatal(err)
		}
	}

	// create/read the file for logging
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "\033[34m"+"INFO: "+"\033[0m", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "\033[31m"+"ERROR :"+"\033[0m", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "\033[33m"+"WARNING: "+"\033[0m", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	InfoLogger.Println("Starting the api...")

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	http.ListenAndServe(":9999", nil)
}

func welcomeHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "welcome to my api\n")
}
