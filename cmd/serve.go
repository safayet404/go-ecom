package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
	"strconv"
)

func Serve() {

	cnf := config.GetConfig()
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
