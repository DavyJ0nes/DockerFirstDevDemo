package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestDataHandler checks that the data handler retrns 200
func TestDataHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(dataHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("Error Getting Data: %s", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected 200 | Got: %v", res.StatusCode)
	}
}

// TestHealthHandler checks that the data handler retrns 200
func TestHealthHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(healthHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("Error Getting Health: %s", err)
	}

	if res.StatusCode != 200 {
		t.Errorf("Expected 200 | Got: %v", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error Parsing Body: %s", err)
	}

	healthData := health{}
	err = json.Unmarshal([]byte(body), &healthData)
	if err != nil {
		t.Errorf("Error Parsing Body in JSON: %s", err)
	}

	if healthData.Status != 1 {
		t.Errorf("Expected 1 | Got: %v", healthData.Status)
	}

	if healthData.Message != "A OK" {
		t.Errorf("Expected 'A OK' | Got: %v", healthData.Message)
	}
}
