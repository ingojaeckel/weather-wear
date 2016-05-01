package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
)

var memcacheClient *memcache.Client

func initializeCache() {
	if !cacheEnabled {
		return
	}
	host := os.Getenv("MEMCACHE_PORT_11211_TCP_ADDR")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("MEMCACHE_PORT_11211_TCP_PORT")
	if port == "" {
		port = "11211"
	}
	memcacheClient = memcache.New(fmt.Sprintf("%s:%s", host, port))
}

func cachePut(key string, val string, expirationSeconds int32) error {
	if !cacheEnabled {
		return nil
	}
	initial := &memcache.Item{
		Key:        key,
		Value:      []byte(val),
		Expiration: expirationSeconds,
	}
	return memcacheClient.Add(initial)
}

func cacheGet(key string) (string, error) {
	if !cacheEnabled {
		return "", errors.New("Cache disabled")
	}
	item, err := memcacheClient.Get(key)
	if err != nil {
		fmt.Printf("cache error: %s\n", err.Error())
		return "", err
	}
	return string(item.Value), nil

}
