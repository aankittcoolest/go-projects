package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
    logger *log.Logger
)

func hello(w http.ResponseWriter, req *http.Request) {
	message := "Hello from QA"
	logger.Println(message)
	fmt.Fprintf(w, message + "\n")
}

func main() {
	// set up port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	// set up logging
    file, err := os.OpenFile("./logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

	hostname, _ := os.Hostname()
	logger = log.New(file, "INFO: " + hostname + " ", log.Ldate|log.Ltime|log.Lshortfile,)
    logger.SetOutput(file)

	// say hello
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":" + port, nil)
}