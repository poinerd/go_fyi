package main

import(
	"net/http"
	"fmt"
	"encoding/json"
)


type product struct{
	ProductName string `json:"ProductName"`
	ProductPrice float64 `json:"ProductPrice"`
}

var Products []product


func getUrlParameter(w http.ResponseWriter, r *http.Request){
	stuffToGet := r.URL.Query().Get("product-name")
	return stuffToGet
}

func getInventory(w http.ResponseWriter, r *http.Request){
     json.NewEncoder(w).Encode(Products)
}

func deleteAProduct(w http.ResponseWriter, r *http.Request){
   productToDelete := getUrlParameter()

   for i, product := range Products{
	if product.ProductName == productToDelete{
		Products = append(Products[:1], Products[i:]...)
		break
	}
   }
   json.NewEncoder(w).Encode(Products)
}

func editAProduct(w http.ResponseWriter, r *http.Request){
	productToEdit := getUrlParameter()

	var editedProduct product
	json.NewDecoder(r).Decode(&editedProduct)

	for i, product := range Products{
		if product.ProductName == productToEdit{
			Products[i] = editedProduct

			break
		}

	}

}

func main(){
	http.ListenAndServe(":8000", nil)

}