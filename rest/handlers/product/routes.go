package product

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /safayet", manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProduct), h.middlewares.AuthJwt))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(h.GetProductByID)))
	mux.Handle("PATCH /products/{productId}", manager.With(http.HandlerFunc(h.UpdateProduct), h.middlewares.AuthJwt))
	mux.Handle("DELETE /products/{productId}", manager.With(http.HandlerFunc(h.DeleteProduct), h.middlewares.AuthJwt))

}
