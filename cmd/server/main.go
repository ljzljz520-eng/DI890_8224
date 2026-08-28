package main

import (
	"log"
	"memorial/api"
	"memorial/service"
	"memorial/store"
	"net/http"
)

func main() {
	s, e := store.Open("memorial.db")
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	log.Fatal(http.ListenAndServe(":8080", api.New(service.New(s))))
}
