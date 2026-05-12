package main

import (
	"net/http"
)

func hiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello! You requested: /hello"))
}

//func main() {
//	http.HandleFunc("/hello", hiHandler)
//	fmt.Println("Server running on http://localhost:8080")
//
//	server := &http.Server{Addr: ":8080", Handler: handler}
//
//	server.ListenAndServe()
//}
