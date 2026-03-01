package handlers

import (
	"ecommerce/product"
	"ecommerce/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	util.SendData(w, product.ProductList, 200)

}
