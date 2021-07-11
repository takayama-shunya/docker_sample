package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.handleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("received request")
		fmt.Fprintf(w, "Hello Docker!!")
	})

	log.Println("start serve")
	serve := &http.Server{Addr: ":8080"}
	if err := server.ListenAndSere(); err != nil {
		log.Println(err)
	}
}