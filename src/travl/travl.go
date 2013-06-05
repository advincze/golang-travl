package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var species = flag.String("species", "gopher", "the species we are studying")

func main() {

	flag.Parse()

	r := mux.NewRouter()

	r.HandleFunc("/{type}", wrapHandler(createObject)).Methods("POST")
	r.HandleFunc("/{type}/{id}", wrapHandler(deleteObject)).Methods("DELETE")
	r.HandleFunc("/{type}/{id}/_av", wrapHandler(defineAvailability)).Methods("PUT")
	r.HandleFunc("/{type}/{id}/_av", wrapHandler(retrieveAvailability)).Methods("GET")
	r.HandleFunc("/{type}/{id}/_ev", wrapHandler(addEvent)).Methods("PUT")

	http.Handle("/", r)

	http.ListenAndServe(":1982", nil)

}

func wrapHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Printf("wrap with %v\n", vars)
		fn(w, r)
	}
}

func createObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := vars["type"]
	body, _ := ioutil.ReadAll(r.Body)

	type Message struct {
		Id         string
		Resolution string
	}
	fmt.Printf("body: %s\n", body)
	var v *Message
	err := json.Unmarshal(body, &v)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "create  %v %v\n", t, v)

}

func deleteObject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete res\n")
}

func defineAvailability(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "defineAvailability\n")
}

func retrieveAvailability(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "retrieveAvailability\n")
}

func addEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "addEvent\n")
}
