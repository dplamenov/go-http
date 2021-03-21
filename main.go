package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "templates/home.html")
}

func about(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "templates/about.html")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	addr := ":8080"
	fmt.Println("Listen on", addr)
	http.ListenAndServe(addr, nil)
}