package main

import (
	"encoding/json"
	"net/http"
)

type Item struct {
	ID       int    `json:"id"` // тег
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

// массив из объектов в Golang
var items = []Item{
	{ID: 1, Name: "Apple", Price: 100, Category: "Fruit"},
	{ID: 2, Name: "Banana", Price: 50, Category: "Fruit"},
	{ID: 3, Name: "Bread", Price: 200, Category: "Bakery"},
	{ID: 4, Name: "Milk", Price: 150, Category: "Dairy"},
}

func getItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // content type = JSON
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

func getOneItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items[0])
}

// GIN, Gorilla Mux, Chi, Echo
//func main() {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/all", getItemsHandler)
//	mux.HandleFunc("/one", getOneItemHandler)
//
//	server := &http.Server{Addr: ":8080", Handler: mux}
//
//	fmt.Println("Server running on http://localhost:8080")
//	if err := server.ListenAndServe(); err != nil {
//		return
//	}
//}
