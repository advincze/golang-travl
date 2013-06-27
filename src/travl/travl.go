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

		ob.Ba.Set(m.From, m.To, m.Value)

		fmt.Fprintf(w, "defineAvailability, type: %v , id: %v , %v, %v n", t, id, ob, m)
	}
}

func retrieveAvailability(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	from, _ := ParseTimeWithMultipleLayouts(q.Get("from"), time.RFC3339, "2006-01-02T15:04", "2006-01-02 15:04", "2006-01-02")
	to, _ := ParseTimeWithMultipleLayouts(q.Get("to"), time.RFC3339, "2006-01-02T15:04", "2006-01-02 15:04", "2006-01-02")
	resolutionStr := q.Get("resolution")
	res := av.ParseTimeResolution(resolutionStr)

	vars := mux.Vars(r)
	t := vars["type"]
	id := vars["id"]
	_, ob := av.GetObjectTypeAndObject(t, id)
	bv := ob.Ba.Get(from, to, res)
	fmt.Println("travl rA", ob.Ba)
	bb, _ := json.Marshal(bv)
	fmt.Fprint(w, string(bb))
	//fmt.Fprintf(w, "retrieveAvailability, %s, %v, %v, %s, %s, %v \n", res, from, to, t, id, bv)
}

func ParseTimeWithMultipleLayouts(s string, layouts ...string) (time.Time, bool) {
	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t, true
		}
	}
	return time.Now(), false
}

func addEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "addEvent\n")
}
