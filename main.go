package main

import (
	"fmt"
	"net/http"

	"github.com/subodhqss/go-training/controllers"
)

func main() {
	StartServer()
	fmt.Println("Hello Trainee !")
}

func StartServer() {
	r := controllers.NewRouter()
	server := &http.Server{
		Addr:    ":4000",
		Handler: r,
	}
	fmt.Println("Server is running at ", server.Addr)
	server.ListenAndServe()
}
