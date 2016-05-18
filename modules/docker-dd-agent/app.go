package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const metadataURL = "http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/access-configs/0/external-ip"

var internalAddress = getInternalAddress()

func getInternalAddress() string {
	req, err := http.NewRequest("GET", metadataURL, nil)
	req.Header.Add("Metadata-Flavor", "Google")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err.Error()
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err.Error()
	}

	return string(body)
}

func getInternalAddressResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, internalAddress)
}

func main() {
	http.HandleFunc("/", getInternalAddressResponse)
	fmt.Println("Running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
