package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Product struct to represent product data
type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float32 `json:"price"`
}

// Slice to store products
var products []Product

// Handler to retrieve all products
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

// Handler to add a new product
func addProductHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var newProduct Product
    err := json.NewDecoder(r.Body).Decode(&newProduct)
    if err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    // Assign a unique ID to the new product
    newProduct.ID = len(products) + 1
    products = append(products, newProduct)

    w.WriteHeader(http.StatusCreated)
    fmt.Fprintf(w, "Product added with ID: %d", newProduct.ID)
}

// Main function to start the server and set up routes
func main() {
    http.HandleFunc("/products", getProductsHandler)
    http.HandleFunc("/addProduct", addProductHandler)

    fmt.Println("Product Service running on port 8081")
    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}