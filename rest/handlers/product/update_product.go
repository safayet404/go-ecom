package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("productId")
	id, err := strconv.Atoi(productID)

	if err != nil {
		fmt.Println("error in post product", err)
		http.Error(w, "Please give me a valid json", 400)
		return
	}

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("error in update product", err)
		http.Error(w, "Please give me a valid json", 400)
		return
	}

	newProduct.ID = id

	database.Update(newProduct)

	util.SendData(w, "Successfully updated product", 201)

}
