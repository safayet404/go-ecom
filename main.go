package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)
	fmt.Fprintln(w, "hello world")
}

func aboutHadnler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)

	fmt.Fprintln(w, "hello this is afayet")
}

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		http.Error(w, "Please give me GET request", 400)

		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)

}

var productList []Product

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/about", aboutHadnler)
	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server running on Server : 3000")

	err := http.ListenAndServe(":3000", mux)

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
