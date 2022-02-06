package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const port = "127.0.0.1:8080"
	router := mux.NewRouter()
	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Server is Up")
	})

	router.HandleFunc("/Patient", getPatient).Methods("GET")
	router.HandleFunc("/Patient", addPatient).Methods("POST")

	log.Println("Serving Listing at ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
