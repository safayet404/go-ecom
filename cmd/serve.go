package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	manager := middleware.NewManager()
	mux := http.NewServeMux()

	manager.Use(middleware.CorsWithPreflight, middleware.Logger)

	initRoutes(mux, manager)

	fmt.Println("Server running on Server : 65535")

	wrapperMux := manager.WrapMux(mux)

	err := http.ListenAndServe(":65535", wrapperMux)

	if err != nil {
		fmt.Println("error in staring the server,try again", err)
	}
}
