package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/alicebob/miniredis"
)

func TestApiHandler(t *testing.T) {
	redisServer, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer redisServer.Close()

	os.Setenv("REDIS_HOST", redisServer.Addr())

	ts := httptest.NewServer(http.HandlerFunc(apiHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("Error Getting API: %s", err)
	}

	// test if request Successful
	if res.StatusCode != 200 {
		t.Errorf("Expected 200 | Got %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("Error Response Parsing Body: %s", err)
	}

	resBody := data{}
	err = json.Unmarshal([]byte(body), &resBody)

	if err != nil {
		t.Errorf("Error Response Parsing Body in JSON: %s", err)
	}
	if resBody.Hitcount != 1 {
		t.Errorf("Expected: 1 | Got: %d", resBody.Hitcount)
	}
}

func TestIndexHandler(t *testing.T) {
	redisServer, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer redisServer.Close()

	os.Setenv("REDIS_HOST", redisServer.Addr())

	ts := httptest.NewServer(http.HandlerFunc(indexHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("Error Getting Index: %s", err)
	}

	// test if request Successful
	if res.StatusCode != 200 {
		t.Errorf("Expected 200 | Got %d", res.StatusCode)
	}
}
