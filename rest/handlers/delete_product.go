package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")

	id, err := strconv.Atoi(productId)

	if err != nil {
		fmt.Println(err)
	}

	database.Delete(id)

	util.SendData(w, "Product deleted", 200)
}
