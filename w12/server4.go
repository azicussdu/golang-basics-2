package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{ID: 1, Name: "Laptop", Price: 2300.5},
	{ID: 2, Name: "IPhone", Price: 1300.5},
	{ID: 3, Name: "Samsung", Price: 3300.5},
}

func main() {
	mux := http.NewServeMux() // golang 1.22

	mux.HandleFunc("GET /products", getAllProducts)
	mux.HandleFunc("GET /products/{id}", getProductByID)
	mux.HandleFunc("POST /products", createProduct)
	mux.HandleFunc("PUT /products/{id}", updateByID)
	mux.HandleFunc("DELETE /products/{id}", deleteProductByID)
	//localhost:8080/products/2   [id="2"]

	server := &http.Server{Addr: ":8080", Handler: mux}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("server start error")
		return
	}
}

func updateByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr) // "2" -> 2
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedProduct Product
	err = json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, prd := range products {
		if prd.ID == id {
			updatedProduct.ID = prd.ID
			products[i] = updatedProduct
			err = json.NewEncoder(w).Encode(updatedProduct)
			if err != nil {
				return
			}
			return
		}
	}

	http.Error(w, "product not found", http.StatusNotFound)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	newProduct.ID = products[len(products)-1].ID + 1
	products = append(products, newProduct)

	w.WriteHeader(http.StatusCreated)
	//w.Write([]byte(`{"message":"product created successfully"}`))
	err = json.NewEncoder(w).Encode(newProduct)
	if err != nil {
		return
	}
}

func getProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr) // "2" -> 2
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for _, prd := range products {
		if prd.ID == id {
			json.NewEncoder(w).Encode(prd)
			return
		}
	}

	http.Error(w, "product not found", http.StatusNotFound)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, "json encode error", http.StatusInternalServerError)
		return
	}
}

func deleteProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr) // "2" -> 2
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, prd := range products {
		if prd.ID == id {
			products = append(products[:i], products[i+1:]...)
			w.Write([]byte(`{"message":"product deleted"}`))
			return
		}
	}

	http.Error(w, "product not found", http.StatusNotFound)
}
