package rest

import (
	"ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /safayet", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct)))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(handlers.GetProductByID)))
	mux.Handle("PATCH /products/{productId}", manager.With(http.HandlerFunc(handlers.UpdateProduct)))
	mux.Handle("DELETE /products/{productId}", manager.With(http.HandlerFunc(handlers.DeleteProduct)))

	// Users

	mux.Handle("GET /users", manager.With(http.HandlerFunc(handlers.GetUsers)))

	mux.Handle("POST /users", manager.With(http.HandlerFunc(handlers.CreateUsers)))
}
