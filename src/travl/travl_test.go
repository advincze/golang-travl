package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldCreateObject(t *testing.T) {

	ts := httptest.NewServer(getHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/obj")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("XXX%sYYY, %s", body, ts.URL)

}
