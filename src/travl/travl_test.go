package main

import (
	"bytes"
	"encoding/json"
	"github.com/steveyen/gkvlite"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestDB(t *testing.T) {
	fname := "myfile.db"
	f, _ := os.Create(fname)
	defer os.Remove(fname)
	s, _ := gkvlite.NewStore(f)
	c := s.SetCollection("cars", nil)

	c.Set([]byte("tesla"), []byte("$$$"))
	c.Set([]byte("mercedes"), []byte("$$"))
	c.Set([]byte("bmw"), []byte("$"))

	s.Flush()
	f.Sync()
}

func TestShouldCreateObjectWithID(t *testing.T) {
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
	var response struct {
		Id string `json:"id"`
	}
	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		t.Errorf("could not parse json document")
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("statuscode should be 'Created', was %v\n", res.StatusCode)
	}
	if response.Id != "8" {
		t.Errorf("id of 'Created', should be 8, was %v\n", response.Id)
	}
}

func TestShouldCreateObjectWithoutID(t *testing.T) {
	ts := httptest.NewServer(createRouter())
	defer ts.Close()

	res, err := http.Post(ts.URL+"/obj", "appication/json", nil)
	if err != nil {
		log.Fatal(err)
	}
	var response struct {
		Id string `json:"id"`
	}
	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		t.Errorf("could not parse json document")
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("statuscode should be 'Created', was %v\n", res.StatusCode)
	}
	if response.Id == "" {
		t.Errorf("id of 'Created', should not be empty")
	}
}

func TestShouldDefineAvailability(t *testing.T) {
	ts := httptest.NewServer(createRouter())
	defer ts.Close()
	createmsg, _ := json.Marshal(struct {
		From  time.Time `json:"from"`
		To    time.Time `json:"to"`
		Value byte      `json:"value"`
	}{time.Now(), time.Now(), 1})

	req, _ := http.NewRequest("PUT", ts.URL+"/obj/5/_av", bytes.NewBuffer(createmsg))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("statuscode should be OK, was %v\n", res.StatusCode)
	}

}
