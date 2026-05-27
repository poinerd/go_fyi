package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type product struct {
	ProductName string  `json:"ProductName"` // fields must start in capitals in Go
	Price       float64 `json:"Price"`
}

var inventory = []product{
	{ProductName: "Colgate Paste", Price: 45.00},
}

// controllers

func addProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct product
	json.NewDecoder(r.Body).Decode(&newProduct)
	inventory = append(inventory, newProduct)
	json.NewEncoder(w).Encode(inventory)
}


func getInventory(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed) // Sends a 405 error back
        json.NewEncoder(w).Encode(map[string]string{"error": "Only GET requests are allowed on this endpoint"})
        return
    }
}

func deleteAProduct(w http.ResponseWriter, r *http.Request) {
	productToDelete := r.URL.Query().Get("name")

	for i, product := range inventory {
		if product.ProductName == productToDelete {
			inventory = append(inventory[:i], inventory[i+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(inventory)
}

func deleteAllProduct(w http.ResponseWriter, r *http.Request) {
	inventory = []product{}
	json.NewEncoder(w).Encode(inventory)
}

func editAProduct(w http.ResponseWriter, r *http.Request) {
	productToEdit := r.URL.Query().Get("product")

	var editedProduct product
	json.NewDecoder(r.Body).Decode(&editedProduct)
	
	for i, product := range inventory {
		if product.ProductName == productToEdit {
			inventory[i] = editedProduct
			break
		}
	}
	json.NewEncoder(w).Encode(inventory)
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allows browser contexts to pull data from this server
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Catches preflight OPTIONS check sent automatically by browsers during PUT/DELETE operations
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// routes

func main() {
	// POST, GET , DELETE, PUT METHODS
	http.HandleFunc("/addproduct", addProduct)
	http.HandleFunc("/getinventory", getInventory)
	http.HandleFunc("/deleteaproduct", deleteAProduct)
	http.HandleFunc("/deleteallproduct", deleteAllProduct)
	http.HandleFunc("/editaproduct", editAProduct)

	fmt.Println("server is running on port 8000")

	// Intercept default multiplexer and route traffic through middleware rules
	defaultRouter := http.DefaultServeMux
	corsProtectedRouter := enableCORS(defaultRouter)

	// Fire up the listener passing the wrapped engine 
	http.ListenAndServe(":8000", corsProtectedRouter)
}