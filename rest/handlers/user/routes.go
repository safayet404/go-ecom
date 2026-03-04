package user

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	// Users

	mux.Handle("GET /users", manager.With(http.HandlerFunc(h.GetUsers), middleware.AuthJwt))

	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUsers), middleware.AuthJwt))
	mux.Handle("POST /login", manager.With(http.HandlerFunc(h.LoginUser)))
}
