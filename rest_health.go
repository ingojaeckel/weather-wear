package main

import (
	"fmt"
	"net/http"

	"github.com/bradfitz/gomemcache/memcache"
)

func getHealth(w http.ResponseWriter, r *http.Request) {
	if cacheEnabled {
		item := &memcache.Item{
			Key:   "health_check",
			Value: []byte{},
		}
		// check memcache connectivity
		if err := memcacheClient.Set(item); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "memcache connectivity: %s", err.Error())
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "ok")
}
