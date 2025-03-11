package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Root Page"))
	})

	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Next Page"))
	})

	log.Fatal(http.ListenAndServe(":3030", nil))
}
