package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}

func defaultHandler(w http.ResponseWriter, req *http.Request) {
}
