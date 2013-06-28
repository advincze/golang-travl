package main

import (
	"bytes"
	"encoding/json"
	"github.com/steveyen/gkvlite"
	"io/ioutil"
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
		t.Errorf("id of 'Created', should be 8, was %v\n", response)
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

func createDefineAvMessage(from, to time.Time, value byte) []byte {
	createmsg, _ := json.Marshal(struct {
		From  time.Time `json:"from"`
		To    time.Time `json:"to"`
		Value byte      `json:"value"`
	}{from, to, value})
	return createmsg
}

func callPUT(url string, msg []byte) *http.Response {
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(msg))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func TestShouldDefineAvailability(t *testing.T) {
	ts := httptest.NewServer(createRouter())
	defer ts.Close()

	createmsg := createDefineAvMessage(time.Now(), time.Now().Add(25*time.Hour), 1)
	res := callPUT(ts.URL+"/obj/5/_av", createmsg)
	if res.StatusCode != http.StatusOK {
		t.Errorf("statuscode should be OK, was %v\n", res.StatusCode)
	}

}

func TestShouldRetrieveAvailability(t *testing.T) {
	ts := httptest.NewServer(createRouter())
	defer ts.Close()
	createmsg := createDefineAvMessage(time.Now(), time.Now().Add(45*time.Hour), 1)
	callPUT(ts.URL+"/obj/5/_av", createmsg)
	res, err := http.Get(ts.URL + "/obj/5/_av?from=2013-06-28&to=2013-06-30&resolution=day")
	if err != nil {
		t.Error(err.Error())
	}
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	println("contents", string(contents))
}
