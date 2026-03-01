package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /create-products", http.HandlerFunc(createProduct))

	fmt.Println("Server running on Server : 3000")

	globalRouter := globalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		fmt.Println("error in staring the server", err)
	}

}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "This is orage fruit",
		Price:       123.45,
		ImgUrl:      "dole.com/sites/default/files/mdedia/2025-01/orange.png",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "This is orage fruit",
		Price:       123.45,
		ImgUrl:      "dole.com/sites/default/files/mdedia/2025-01/orange.png",
	}
	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "This is orage fruit",
		Price:       123.45,
		ImgUrl:      "dole.com/sites/default/files/mdedia/2025-01/orange.png",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
}
