package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

var (
	version = "No Version Provided"
	date    = ""
	gitHash = ""
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/data", dataHandler)
	mux.HandleFunc("/health", healthHandler)
	log.Println("Starting Go API Server")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

type data struct {
	Name         string
	RandomString string
	Version      string
}

// dataHandler returns data for GET requests to /v1/data
func dataHandler(w http.ResponseWriter, req *http.Request) {
	requestLogger(req)
	versionString := fmt.Sprintf("GoAPI - %s (%s) %s", version, gitHash, date)

	data := data{
		"Go API",
		generateRandomString(),
		versionString,
	}

	js, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

type health struct {
	Status  int
	Message string
}

// healthHandler is used to check the health of the service.
// Health endpoints can either look for a 200 request or specific data
// Can also be extended in the future to be more dynamic. Showing performance data.
func healthHandler(w http.ResponseWriter, req *http.Request) {
	healthData := health{
		Status:  1,
		Message: "A OK",
	}

	status, err := json.Marshal(healthData)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(status)
}

// requestLogger is a helper function that prints request logging information.
func requestLogger(req *http.Request) {
	log.Printf("%s | %s => %s", req.Method, req.RemoteAddr, req.URL.Path)
}

// generateRandomString generates a randomString... useless...comment...
func generateRandomString() string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 64)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
