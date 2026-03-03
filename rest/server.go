package rest

import (
	"ecommerce/config"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"strconv"
)

func Start(cnf config.Config) {

	manager := middleware.NewManager()
	mux := http.NewServeMux()

	manager.Use(middleware.CorsWithPreflight, middleware.Logger)

	initRoutes(mux, manager)
	address := ":" + strconv.Itoa(int(cnf.HttpPort))

	fmt.Println("Server running on ", address)

	wrapperMux := manager.WrapMux(mux)
	err := http.ListenAndServe(address, wrapperMux)

	if err != nil {
		fmt.Println("error in staring the server,try again", err)
	}
}
