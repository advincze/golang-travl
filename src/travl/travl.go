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

	http.Handle("/", getHandler())

	http.ListenAndServe(":1982", nil)

}

func getHandler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/{type}", wrapHandler(createObject)).Methods("POST")
	r.HandleFunc("/{type}/{id}", wrapHandler(deleteObject)).Methods("DELETE")
	r.HandleFunc("/{type}/{id}/_av", wrapHandler(defineAvailability)).Methods("PUT")
	r.HandleFunc("/{type}/{id}/_av", wrapHandler(retrieveAvailability)).Methods("GET")
	r.HandleFunc("/{type}/{id}/_ev", wrapHandler(addEvent)).Methods("PUT")
	r.HandleFunc("/{type}", wrapHandler(infoHandler)).Methods("GET")
	return r
}

func wrapHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Printf("wrap with %v\n", vars)
		fn(w, r)
	}
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "info")
}

func createObject(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	type Message struct {
		Id         string `json:"id"`
		Resolution string `json:"resolution"`
	}

	var v *Message
	err := json.Unmarshal(body, &v)
	if err != nil {
		http.Error(w, "could not parse json document", http.StatusInternalServerError)
	}

	// save v
	v.Resolution = ""

	type Resp struct {
		Id string `json:"id"`
	}

	resp := &Resp{Id: v.Id}

	bytes, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", bytes)

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
