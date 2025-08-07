package main

import (
	"fmt"
	"net/http"
)

func main() {
	//config := configs.NewConfig()

	router := http.NewServeMux()

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Printf("Server start on %s port", server.Addr)
	server.ListenAndServe()

}
