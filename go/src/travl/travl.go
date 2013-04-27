package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/res", wrapHandler(createHandler)).Methods("PUT")
	r.HandleFunc("/res/", wrapHandler(createHandler)).Methods("PUT")
	r.HandleFunc("/res/{id}", wrapHandler(readHandler)).Methods("GET")
	r.HandleFunc("/res/{id}/", wrapHandler(readHandler)).Methods("GET")
	r.HandleFunc("/res/{id}", wrapHandler(updateHandler)).Methods("PUT")
	r.HandleFunc("/res/{id}/", wrapHandler(updateHandler)).Methods("PUT")
	r.HandleFunc("/res/{id}", wrapHandler(deleteHandler)).Methods("DELETE")
	r.HandleFunc("/res/{id}/", wrapHandler(deleteHandler)).Methods("DELETE")

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

func wrapHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Printf("wrap with %v", vars)
		fn(w, r)
	}
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "create res\n")
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	fmt.Fprintf(w, "Hi there, I love %s! \n", id)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update res\n")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete res\n")
}
