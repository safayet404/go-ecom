package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)
	fmt.Fprintln(w, "hello world")
}

func aboutHadnler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.Path)

	fmt.Fprintln(w, "hello this is afayet")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/about", aboutHadnler)

	fmt.Println("Server running on Server : 3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("error in staring the server", err)
	}

}
