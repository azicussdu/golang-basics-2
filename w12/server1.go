package main

import (
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello! You requested: /hello"))
}

//func main() {
//	http.HandleFunc("/hello", helloHandler)
//	fmt.Println("Server running on http://localhost:8080")
//	http.ListenAndServe(":8080", nil)
//}
