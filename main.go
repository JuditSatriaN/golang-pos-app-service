package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/golang_pos_app_service/handler/productinfo"
	"github.com/gorilla/mux"
)

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	// products handler
	myRouter.HandleFunc("/products", productinfo.GetProducts)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Starting App Service")
	handleRequests()
}
