package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")

	id, err := strconv.Atoi(productId)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "please give me a valid product id", 400)
		return
	}

	err = h.productRepo.Delete(id)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Product deleted", 200)
}
