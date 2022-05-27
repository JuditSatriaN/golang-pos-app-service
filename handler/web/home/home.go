package home

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/golang_pos_app_service/entity/constants"
)

func WebHomeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(constants.HeaderViewHtml, constants.LoadingViewHtml, constants.SideBarViewHtml, constants.FooterViewHtml, constants.HomeViewHtml)
	if err != nil {
		fmt.Fprintf(w, "Home Page Error, Please Try Again Later, Err : %v", err)
	}

	err = t.ExecuteTemplate(w, "home.html", constants.WebData{
		Title:        constants.WebHomeTitle,
		BaseURL:      constants.BaseURL,
		TemplateURL:  constants.TemplateUrl,
		LinkPageList: constants.LinkPageList,
	})
	if err != nil {
		fmt.Fprintf(w, "Home Page Error, Please Try Again Later, Err : %v", err)
	}
}
