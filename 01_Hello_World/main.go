package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello you have requested:%s\n", r.URL.Path)
	})
	fmt.Println("We are listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
