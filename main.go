package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	personaliaSvc "github.com/golang_pos_app_service/handler/services/personalia"
	productinfoSvc "github.com/golang_pos_app_service/handler/services/productinfo"

	homeWeb "github.com/golang_pos_app_service/handler/web/home"
	loginWeb "github.com/golang_pos_app_service/handler/web/login"
	personaliaWeb "github.com/golang_pos_app_service/handler/web/personalia"
)

func handleRequests() {
	// creates a new instance of a mux router

	myRouter := mux.NewRouter().StrictSlash(true)

	// Web handler
	var static string
	flag.StringVar(&static, "static", "./static", "")
	flag.Parse()
	myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(static))))
	myRouter.HandleFunc("/", loginWeb.WebLoginHandler)
	myRouter.HandleFunc("/home", homeWeb.WebHomeHandler)
	myRouter.HandleFunc("/user", personaliaWeb.WebUserHandler)
	myRouter.HandleFunc("/member", personaliaWeb.WebMemberHandler)

	// Service handler
	// Service - products
	myRouter.HandleFunc("/svc/products", productinfoSvc.GetProducts)

	// Service - users
	myRouter.HandleFunc("/svc/users", personaliaSvc.ServicesGetUsers)
	myRouter.HandleFunc("/svc/user/insert", personaliaSvc.ServicesInsertUser)
	myRouter.HandleFunc("/svc/user/update", personaliaSvc.ServicesUpdateUser)
	myRouter.HandleFunc("/svc/user/delete", personaliaSvc.ServicesDeleteUser)
	myRouter.HandleFunc("/svc/user/upsert", personaliaSvc.ServicesUpsertUser)
	// Service - members
	myRouter.HandleFunc("/svc/members", personaliaSvc.ServicesGetMembers)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Starting App Service")
	handleRequests()
}
