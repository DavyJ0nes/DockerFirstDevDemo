package main

import (
	"encoding/json"
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
	mux.HandleFunc("/v1/data", apiHandler)
	mux.HandleFunc("/v1/version", versionHandler)
	log.Println("Starting API Server")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

// data
type data struct {
	Name         string
	RandomString string
}

// apiHandler returns data for GET requests to /v1/data
func apiHandler(w http.ResponseWriter, req *http.Request) {
	requestLogger(req)
	// versionString := fmt.Sprintf("GoAPI - %s (%s) %s", version, gitHash, date)

	apiData := data{
		"Go API",
		generateRandomString(),
	}

	js, err := json.Marshal(apiData)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// version
type versionInfo struct {
	Version   string
	GitHash   string
	BuildDate string
}

// versionHandler
func versionHandler(w http.ResponseWriter, req *http.Request) {
	requestLogger(req)
	versionInfo := versionInfo{
		version,
		gitHash,
		date,
	}

	js, err := json.Marshal(versionInfo)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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
