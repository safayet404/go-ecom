package rest

import (
	"ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /safayet", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct), middleware.AuthJwt))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(handlers.GetProductByID)))
	mux.Handle("PATCH /products/{productId}", manager.With(http.HandlerFunc(handlers.UpdateProduct), middleware.AuthJwt))
	mux.Handle("DELETE /products/{productId}", manager.With(http.HandlerFunc(handlers.DeleteProduct), middleware.AuthJwt))

	// Users

	mux.Handle("GET /users", manager.With(http.HandlerFunc(handlers.GetUsers), middleware.AuthJwt))

	mux.Handle("POST /users", manager.With(http.HandlerFunc(handlers.CreateUsers), middleware.AuthJwt))
	mux.Handle("POST /login", manager.With(http.HandlerFunc(handlers.LoginUser)))
}
