package main

import (
	"fmt"
	"log"
	"net/http"
)

func apiAHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	for k, v := range r.Header {
		log.Println("key:", k, "value", v)
	}
	log.Println("-------------------------------------------------------")
	response := "Hello from API A!"
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, response)
}

func apiBHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	for k, v := range r.Header {
		log.Println("key:", k, "value", v)
	}
	log.Println("-------------------------------------------------------")

	response := "Hello from API B!"
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, response)
}

func main() {
	http.HandleFunc("/api/a", apiAHandler)
	http.HandleFunc("/api/b", apiBHandler)

	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
