package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/mediocregopher/radix.v2/redis"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	version = "No Version Provided"
	date    = ""
	gitHash = ""
)

var requestCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "go_app_request_count_total",
	Help: "Number of requests",
})

func main() {
	prometheus.MustRegister(requestCount)

	mux := http.NewServeMux()
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/v1/data", apiHandler)
	mux.HandleFunc("/v1/version", versionHandler)
	mux.Handle("/metrics", promhttp.Handler())

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

// indexHandler is the root hander for the server
func indexHandler(w http.ResponseWriter, req *http.Request) {
	requestLogger(req)

	versionString := fmt.Sprintf("GoAPI - %s (%s) %s", version, gitHash, date)
	hostname := getHostname()

	redisAddr := os.Getenv("REDIS_HOST")

	client := redisConnect(redisAddr)
	defer client.Close()

	hitCount := increment(client)

	html := []byte(fmt.Sprintf("<h3>Hello from %s</h3>\n<p>Hit Count = %d</p>\n\n<small>%s</small>", hostname, hitCount, versionString))
	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}

// data is simple struct used for JSON output of apiHandler
type data struct {
	Hostname     string `json:"hostname,omitempty"`
	Hitcount     int    `json:"hitcount,omitempty"`
	RandomString string
}

// apiHandler returns JSON
func apiHandler(w http.ResponseWriter, req *http.Request) {
	requestLogger(req)

	hostname := getHostname()
	redisAddr := os.Getenv("REDIS_HOST")
	client := redisConnect(redisAddr)
	defer client.Close()
	hitCount := increment(client)

	data := data{
		hostname,
		hitCount,
		generateRandomString(),
	}

	js, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
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

// redisConnect connects to redis instance and returns client
func redisConnect(redisHost string) *redis.Client {
	client, err := redis.Dial("tcp", redisHost)
	if err != nil {
		log.Panicln("client |", err)
	}
	return client
}

// increment adds 1 to the hit counter in redis
func increment(client *redis.Client) int {
	hits, err := client.Cmd("INCR", "hits").Int()
	if err != nil {
		log.Fatal("hits |", err)
	}
	return hits
}

// requestLogger is a helper function that prints request logging information.
func requestLogger(req *http.Request) {
	requestCount.Add(1)
	log.Printf("%s | %s => %s", req.Method, req.RemoteAddr, req.URL.Path)
}

// getHostname returns the hostname of the node that
// the bainry is running on
func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Println(err)
	}
	return hostname
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
