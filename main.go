package main

import (
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

	http.ListenAndServe(":8090", nil)
}