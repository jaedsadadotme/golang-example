package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Hello Golang")
	})

	log.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
