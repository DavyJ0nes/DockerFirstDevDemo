package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	testURL := flag.String("url", "http://localhost:3000/health", "url to test in format: http://localhost:3000/health")
	flag.Parse()

	_, err := runHealthCheck(*testURL)
	if err != nil {
		fmt.Printf("❗ Healthcheck Failed with Error: '%v'\n", err)
		os.Exit(1)
	}
	fmt.Println("👍 Healthcheck Passed")
}

func runHealthCheck(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		errorString := fmt.Sprintf("Non 200 Status Received. Got: %d", res.StatusCode)
		return "", errors.New(errorString)
	}

	bytes, _ := ioutil.ReadAll(res.Body)
	body := fmt.Sprintf("%s", bytes)

	return body, nil
}
