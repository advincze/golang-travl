package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
	"travl/av"
)

var species = flag.String("species", "gopher", "the species we are studying")

func main() {

	flag.Parse()

	http.Handle("/", createRouter())

	http.ListenAndServe(":1982", nil)

}

func createRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/{type}", createObject).Methods("POST")
	r.HandleFunc("/{type}/{id}", deleteObject).Methods("DELETE")
	r.HandleFunc("/{type}/{id}/_av", defineAvailability).Methods("PUT")
	r.HandleFunc("/{type}/{id}/_av", retrieveAvailability).Methods("GET")
	r.HandleFunc("/{type}/{id}/_ev", addEvent).Methods("PUT")
	r.HandleFunc("/{type}", infoHandler).Methods("GET")
	return r
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "info")
}

func createObject(w http.ResponseWriter, r *http.Request) {

	t := mux.Vars(r)["type"]
	ot := av.GetObjectType(t)
	body, _ := ioutil.ReadAll(r.Body)
	if len(body) != 0 {
		type Message struct {
			Id         string `json:"id"`
			Resolution string `json:"resolution"`
		}

		var v *Message
		err := json.Unmarshal(body, &v)
		if err != nil {
			http.Error(w, "could not parse json document", http.StatusInternalServerError)
		}

		ob := ot.GetObject(v.Id)
		bytes, _ := json.Marshal(ob)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", bytes)
	} else {
		ob := ot.NewObject()
		bytes, _ := json.Marshal(ob)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", bytes)
	}
}

func deleteObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := vars["type"]
	id := vars["id"]
	_, ob := av.GetObjectTypeAndObject(t, id)
	fmt.Fprintf(w, "delete res , type: %v, id: %v , ob: %v \n", t, id, ob)

}

func defineAvailability(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := vars["type"]
	id := vars["id"]
	_, ob := av.GetObjectTypeAndObject(t, id)
	body, _ := ioutil.ReadAll(r.Body)
	if len(body) != 0 {
		type Message struct {
			From  time.Time `json:"from"`
			To    time.Time `json:"to"`
			Value byte      `json:"value"`
		}

		var m *Message
		err := json.Unmarshal(body, &m)
		if err != nil {
			http.Error(w, "could not parse json document", http.StatusInternalServerError)
		}

		ob.Ba.Set(m.From, m.To, m.Value == 1)
		fmt.Fprintf(w, "defineAvailability, type: %v , id: %v , %v, %v n", t, id, ob, m)
	}
}

func retrieveAvailability(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "retrieveAvailability\n")
}

func addEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "addEvent\n")
}
