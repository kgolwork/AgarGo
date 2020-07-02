package main

import (
	"AgarGo/server"
	"log"
	"net/http"
)

func main() {
	s := server.NewServer()
	go s.Listen()

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
