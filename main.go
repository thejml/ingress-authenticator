package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type LogFormat struct {
	Timestamp  time.Time `json:"timestamp"`
	Method     string    `json:"method"`
	RequestURI string    `json:"request"`
	//	Path       string        `json:"internalPath"`
	Duration time.Duration `json:"duration"`
}

// func Logger(inner http.Handler, name string) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		start := time.Now()

// 		//inner.ServeHTTP(w, r)

// 		hit := LogFormat{start, r.Method, r.RequestURI, name, time.Since(start)}
// 		entry, _ := json.Marshal(hit)

// 		log.Print(string(entry))

// 	})
// }

func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("x-api-key") == "deadbeef" {
		fmt.Fprintf(w, "Success")
	} else {
		http.Error(w, "Authorization Failed", http.StatusUnauthorized)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		hit := LogFormat{start, r.Method, r.RequestURI, time.Duration(time.Since(start))}
		entry, _ := json.Marshal(hit)

		log.Print(string(entry))
		handler.ServeHTTP(w, r)
	})
}

func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/auth", authHandler)

	log.Fatal(http.ListenAndServe(":8081", logRequest(http.DefaultServeMux)))

}
