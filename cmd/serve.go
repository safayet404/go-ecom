package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {

	manager := middleware.NewManager()
	mux := http.NewServeMux()

	manager.Use(middleware.Logger, middleware.Hudai)

	initRoutes(mux, manager)

	fmt.Println("Server running on Server : 3000")

	globalRouter := middleware.CorsWithPreflight(mux)

	err := http.ListenAndServe(":3000", globalRouter)

	if err != nil {
		fmt.Println("error in staring the server,try again", err)
	}
}
