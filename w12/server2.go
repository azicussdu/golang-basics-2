package main

import (
	"fmt"
	"net/http"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello! You requested: /hello"))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is about page! Welcome to codee.kz :)")
}

//func main() {
//	mux := http.NewServeMux() // mux -> мультиплексор
//	mux.HandleFunc("/hello", hiHandler)
//	mux.HandleFunc("/about", aboutHandler)
//
//	server := &http.Server{Addr: ":8080", Handler: mux}
//
//	fmt.Println("Server running on http://localhost:8080")
//	if err := server.ListenAndServe(); err != nil {
//		return
//	}
//}
