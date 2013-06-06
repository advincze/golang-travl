package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShouldCreateObject(t *testing.T) {
	ts := httptest.NewServer(createRouter())
	defer ts.Close()

	res, err := http.Post(ts.URL+"/obj", "appication/json", strings.NewReader(`{
		"id"		 : "8",
		"resolution" : "1min"
	}`))
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("statuscode should be 'Created', was %v\n", res.StatusCode)
	}
}
