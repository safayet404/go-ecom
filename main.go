package main

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/product"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProduct))

	fmt.Println("Server running on Server : 3000")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		fmt.Println("error in staring the server", err)
	}

}

func init() {
	prd1 := product.Product{
		ID:          1,
		Title:       "Orange",
		Description: "This is orage fruit",
		Price:       123.45,
		ImgUrl:      "dole.com/sites/default/files/mdedia/2025-01/orange.png",
	}

	prd2 := product.Product{
		ID:          2,
		Title:       "Apple",
		Description: "This is orage fruit",
		Price:       123.45,
		ImgUrl:      "dole.com/sites/default/files/mdedia/2025-01/orange.png",
	}
	prd3 := product.Product{
		ID:          3,
		Title:       "Banana",
		Description: "This is orage fruit",
		Price:       123.45,
		ImgUrl:      "dole.com/sites/default/files/mdedia/2025-01/orange.png",
	}

	product.ProductList = append(product.ProductList, prd1)
	product.ProductList = append(product.ProductList, prd2)
	product.ProductList = append(product.ProductList, prd3)
}
