package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/run/{language}", StreamFile)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:3530",
	}

	log.Fatal(srv.ListenAndServe())
}
