package productinfo

import (
	"fmt"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
