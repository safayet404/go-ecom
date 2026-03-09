package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImgUrl      string `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct ReqCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("error in post product", err)
		http.Error(w, "Please give me a valid json", 400)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       newProduct.Title,
		Description: newProduct.Description,
		Price:       newProduct.Price,
		ImgUrl:      newProduct.ImgUrl,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	util.SendData(w, createdProduct, http.StatusCreated)
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
