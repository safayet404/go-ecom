package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductByID))

	fmt.Println("Server running on Server : 3000")

	globalRouter := global_router.GlobalRouter(mux)

	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		fmt.Println("error in staring the server,try again", err)
	}
}
