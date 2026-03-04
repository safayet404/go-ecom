package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"strconv"
)

type Server struct {
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
	}
}

func (s *Server) Start(cnf config.Config) {

	manager := middleware.NewManager()
	mux := http.NewServeMux()

	manager.Use(middleware.CorsWithPreflight, middleware.Logger)

	// initRoutes(mux, manager)

	s.productHandler.RegisterRoutes(mux, manager)
	s.userHandler.RegisterRoutes(mux, manager)
	address := ":" + strconv.Itoa(int(cnf.HttpPort))

	fmt.Println("Server running on ", address)

	wrapperMux := manager.WrapMux(mux)
	err := http.ListenAndServe(address, wrapperMux)

	if err != nil {
		fmt.Println("error in staring the server,try again", err)
	}
}
