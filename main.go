package main

import (
	"flag"
	"fmt"
	"github.com/golang_pos_app_service/handler/web/login"
	"log"
	"net/http"

	"github.com/golang_pos_app_service/handler/services/productinfo"
	"github.com/golang_pos_app_service/handler/web/home"
	"github.com/gorilla/mux"
)

func handleRequests() {
	// creates a new instance of a mux router

	myRouter := mux.NewRouter().StrictSlash(true)

	// Web handler
	var static string
	flag.StringVar(&static, "static", "./static", "")
	flag.Parse()
	myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(static))))
	myRouter.HandleFunc("/", login.WebLoginHandler)
	myRouter.HandleFunc("/home", home.WebHomeHandler)
	// Service handler
	// Service - products
	myRouter.HandleFunc("/products", productinfo.GetProducts)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Starting App Service")
	handleRequests()
}
