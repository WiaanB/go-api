package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// 5mb log file size limit, adjust to your liking
const FILE_SIZE_LIMIT = 5242880

var (
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	InfoLogger    *log.Logger
	DB            *sql.DB
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
	// load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		ErrorLogger.Println("Failed to read .env file")
		log.Fatal(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, os.Getenv("USERNAME"), os.Getenv("PASSWORD"), os.Getenv("USERNAME"))
	// create a db instance for postgres
	DB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		// should only fail if the information is wrong or the pq driver isn't imported
		ErrorLogger.Println("Failed to create DB instance for postgres")
		log.Fatal(err)
	}

	// test the db connection to see if it works
	err = DB.Ping()
	if err != nil {
		ErrorLogger.Printf("Failed to connect to postgres:%s\n", os.Getenv("USERNAME"))
		log.Fatal(err)
	} else {
		InfoLogger.Printf("Connection established to the %s database\n", os.Getenv("USERNAME"))
	}

	// setup the base tables required for the library example
	initializeTables(DB)

	// setup the routes and listen on port :9999
	mux := http.NewServeMux()
	mux.HandleFunc("/users/", handleUsers)
	mux.HandleFunc("/", welcomeHandler)
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {})

	InfoLogger.Println("Starting the api...")
	defer http.ListenAndServe(":9999", mux)
}

func welcomeHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "welcome to my api\n")
}
