package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("x-api-key") == "deadbeef" {
			fmt.Fprintf(w, "Success")
		} else {
			http.Error(w, "Authorization Failed", 401)
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
