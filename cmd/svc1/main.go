package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the Service 1!")
}

func handleRequests() {
    log.Fatal(http.ListenAndServe(":5001", nil))
    http.HandleFunc("/", homePage)
}

func main() {
    handleRequests()
}
