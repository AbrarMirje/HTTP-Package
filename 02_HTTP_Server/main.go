package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1) Process Dynamic Request
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	// 2) Serving Static Assets
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Listening and Serving on port 9090
	fmt.Println("Listening on port number 9090")
	http.ListenAndServe(":9090", nil)
}
