package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UpdateReqProd struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImgUrl      string `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("productId")
	id, err := strconv.Atoi(productID)

	if err != nil {
		fmt.Println("error in post product", err)
		http.Error(w, "Please give me a valid json", 400)
		return
	}

	var newProduct UpdateReqProd

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("error in update product", err)
		http.Error(w, "Please give me a valid json", 400)
		return
	}

	newProduct.ID = id

	h.productRepo.Update(repo.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImgUrl:      newProduct.ImgUrl,
	})

	util.SendData(w, "Successfully updated product", 201)

}
