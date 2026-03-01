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
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func getProducts(w http.ResponseWriter, r *http.Request) {

	sendData(w, productList, 200)

}

var productList []Product

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func createProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println("error in post product", err)
		http.Error(w, "Please give me a valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1

	productList = append(productList, newProduct)
	sendData(w, newProduct, 201)
}

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

func corsMiddleware(next http.Handler) http.Handler {
	handleCors := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,PATCH,DELETE,OPTIONS,POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)

	}

	handler := http.HandlerFunc(handleCors)
	return handler
}

func globalRouter(mux *http.ServeMux) http.Handler {

	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,PATCH,DELETE,OPTIONS,POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		mux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handleAllReq)
}
