package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestHandler struct{}

func TestHealthCheckOK(t *testing.T) {
	handler := &TestHandler{}
	testSrv := httptest.NewServer(handler)
	defer testSrv.Close()

	res, err := runHealthCheck(testSrv.URL)
	if err != nil {
		t.Log("Testing That Request Was Successful")
		t.Fatal(err)
	}

	expected := `{"string": "testString"}`
	if res != expected {
		t.Log("Testing That Expected String Is Returned")
		t.Errorf("Expected: '%s' || Got: '%s'", expected, res)
	}
}

func TestHealthCheckFail(t *testing.T) {
	handler := &TestHandler{}
	testSrv := httptest.NewServer(handler)
	defer testSrv.Close()

	res, err := runHealthCheck(testSrv.URL)
	if err != nil {
		t.Log("Testing That Request Was Successful")
		t.Fatal(err)
	}

	expected := `{"string": "testString"}`
	if res != expected {
		t.Log("Testing That Expected String Is Returned")
		t.Errorf("Expected: '%s' || Got: '%s'", expected, res)
	}
}

func (h *TestHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	testString := `{"string": "testString"}`
	w.Header().Set("content-type", "applcation/json")
	w.Write([]byte(testString))
}
