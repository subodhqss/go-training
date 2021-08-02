package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/loginc", Loginc)
	http.HandleFunc("/home", Home)
	log.Fatal(http.ListenAndServe(":8500", nil))
}
