package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Potato is a potato
type Potato struct {
	Text    string  `json:"text"`
	History []Entry `json:"history"`
}

// Entry is a historic entry
type Entry struct {
	Node string `json:"node"`
	Text string `json:"text"`
	Desc string `json:"desc"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT environment variable must be set")
	} else {
		fmt.Println("Running on port: " + port)
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/process", Process)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// Process is the main RPC endpoint for this service that will manipulate a passed in string
func Process(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello there and welcome to your service!")
	// TODO: Implement me
}
