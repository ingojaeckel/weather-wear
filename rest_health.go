package main

import (
	"fmt"
	"net/http"
)

func getHealth(w http.ResponseWriter, r *http.Request) {
	// TODO add memcache connectivity
	fmt.Fprint(w, "ok")
}
