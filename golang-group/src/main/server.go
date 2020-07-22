package main

import (
	"group"
	"net/http"
)

func main() {
	http.HandleFunc("/xxx", group.HomeHandler)
	http.ListenAndServe(":8080", nil)
}
