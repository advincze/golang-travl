package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldCreateObject(t *testing.T) {
	ts := httptest.NewServer(createRouter())
	defer ts.Close()

	jsonmsg, _ := json.Marshal(struct {
		Id         string `json:"id"`
		Resolution string `json:"resolution"`
	}{"8", "1min"})

	res, err := http.Post(ts.URL+"/obj", "appication/json", bytes.NewBuffer(jsonmsg))
	if err != nil {
		log.Fatal(err)
	}

	var resp struct {
		Id string `json:"id"`
	}

	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		t.Errorf("could not parse json document")
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("statuscode should be 'Created', was %v\n", res.StatusCode)
	}
}
