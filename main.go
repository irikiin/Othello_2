package main

import (
	"fmt"
	"net/http"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func scriptHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "script.js")
}

func styleHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "style.css")
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/script.js", scriptHandler)
	http.HandleFunc("/style.css", styleHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
