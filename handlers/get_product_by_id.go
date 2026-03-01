package handlers

import (
	"ecommerce/product"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")

	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me avalid product id ", 400)
		return
	}

	for _, product := range product.ProductList {
		if product.ID == id {
			util.SendData(w, product, 200)
			return
		}
	}

	util.SendData(w, "Product pai nai bhai bisshas koren", 404)
}
