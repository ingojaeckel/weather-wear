package main

import (
	"fmt"
	"log"
 	"net"
	"net/http"
	"time"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func connect(w http.ResponseWriter, r *http.Request) {
	if _, err := net.DialTimeout("tcp", "104.197.204.165:8125", 5 * time.Second); err != nil {
		fmt.Fprint(w, err.Error())
		return
	}	
	fmt.Fprint(w, "Connected")
}

func main() {
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/connect", connect)
	fmt.Println("Running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
