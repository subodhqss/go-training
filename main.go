package main

import (
	"fmt"
	"net/http"

	"github.com/subodhqss/go-training/controllers"
	"github.com/subodhqss/go-training/repository"
)

func main() {
	repository.InitDBConnection();
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
