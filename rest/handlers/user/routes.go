package user

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	// Users

	mux.Handle("GET /users", manager.With(http.HandlerFunc(h.GetUsers), h.middlewares.AuthJwt))

	mux.Handle("POST /users", manager.With(http.HandlerFunc(h.CreateUsers), h.middlewares.AuthJwt))
	mux.Handle("POST /login", manager.With(http.HandlerFunc(h.LoginUser)))
}
