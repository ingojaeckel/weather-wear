package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello World!")
}

func main() {
	http.HandleFunc("/hello", helloWorld)
	fmt.Println("Running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
