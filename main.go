package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "./static/index.html")
}
