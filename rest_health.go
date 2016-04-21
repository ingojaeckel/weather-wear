package main

import (
  "fmt"
  "net/http"
)

func GetHealth(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "ok")
}
