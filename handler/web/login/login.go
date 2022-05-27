package login

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/golang_pos_app_service/entity/constants"
)

func WebLoginHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles(constants.LoginViewHtml))
	err := tpl.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "Login Page Error, Please Try Again Later")
	}
}
