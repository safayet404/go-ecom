package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productId")

	id, err := strconv.Atoi(productID)

	if err != nil {
		http.Error(w, "Please give me avalid product id ", 400)
		return
	}

	product, err := h.productRepo.Get(id)

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if product == nil {
		util.SendError(w, http.StatusNotFound, "product not found bhai")
		return
	}

	util.SendData(w, product, 200)
}
