package home

import (
	"fmt"
	"github.com/golang_pos_app_service/entity/constants"
	"github.com/golang_pos_app_service/entity/home"
	"html/template"
	"net/http"
)

func WebHomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(constants.HeaderViewHtml, constants.SideBarViewHtml, constants.FooterViewHtml, constants.HomeViewHtml)
	if err != nil {
		fmt.Fprintf(w, "Home Page Error, Please Try Again Later, Err : %v", err)
	}

	err = t.ExecuteTemplate(w, "home.html", home.WebHomeData{
		Title:       home.WebHomeTitle,
		BaseURL:     constants.BaseURL,
		TemplateURL: constants.TemplateUrl,
	})
	if err != nil {
		fmt.Fprintf(w, "Home Page Error, Please Try Again Later, Err : %v", err)
	}
}
