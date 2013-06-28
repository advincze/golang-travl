package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
	"travl/av"
)

var port = flag.String("port", ":1982", "http port")

func main() {
	flag.Parse()
	http.Handle("/", createRouter())
	http.ListenAndServe(*port, nil)
}

func createRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/{id}/_av", defineAv).Methods("PUT")
	r.HandleFunc("/{id}/_av", retrieveAv).Methods("GET").Queries("from", "", "to", "", "resolution", "")
	r.HandleFunc("/{id}/_ev", addEvent).Methods("PUT")
	r.HandleFunc("/{id}", deleteAv).Methods("DELETE")
	return r
}

func deleteAv(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	av.DeleteBitAv(id)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "")
}

func defineAv(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	bitav := av.FindOrNewBitAv(id)

	var msg *struct {
		From     time.Time     `json:"from"`
		To       time.Time     `json:"to"`
		Duration time.Duration `json:"duration"`
		Value    byte          `json:"value"`
	}

	err := parseBodyJSON(r, &msg)
	if err != nil {
		panic(err)
		http.Error(w, "could not parse json document", http.StatusInternalServerError)
		return
	}

	if msg != nil {
		bitav.Set(msg.From, msg.To, msg.Value)
		bb, _ := json.Marshal(msg)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(bb))
	}
}

func retrieveAv(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	from, _ := parseTime(q.Get("from"))
	to, _ := parseTime(q.Get("to"))
	res := av.ParseTimeResolution(q.Get("resolution"))

	id := mux.Vars(r)["id"]

	bitav := av.FindOrNewBitAv(id)
	bv := bitav.Get(from, to, res)

	bb, _ := json.Marshal(bv)
	fmt.Fprint(w, string(bb))
}

func addEvent(w http.ResponseWriter, r *http.Request) {
	//TODO implement
	fmt.Fprintf(w, "addEvent\n")
}

func parseTime(s string, layouts ...string) (time.Time, error) {
	if len(layouts) == 0 {
		layouts = []string{time.RFC3339, "2006-01-02T15:04", "2006-01-02 15:04", "2006-01-02"}
	}
	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t, nil
		}
	}
	return time.Now(), errors.New("wrong time format")
}

func parseBodyJSON(r *http.Request, data interface{}) error {
	body, _ := ioutil.ReadAll(r.Body)
	if len(body) != 0 {
		return json.Unmarshal(body, &data)
	}
	return nil
}
