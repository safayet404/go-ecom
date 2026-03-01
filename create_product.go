package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func createProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("error in post product", err)
		http.Error(w, "Please give me a valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1

	productList = append(productList, newProduct)
	sendData(w, newProduct, 201)
}
