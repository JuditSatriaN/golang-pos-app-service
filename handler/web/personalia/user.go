package personalia

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/golang_pos_app_service/entity/constants"
)

func WebUserHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(constants.HeaderViewHtml, constants.LoadingViewHtml, constants.SideBarViewHtml, constants.FooterViewHtml, constants.UserViewHtml)
	if err != nil {
		fmt.Fprintf(w, "User Page Error, Please Try Again Later, Err : %v", err)
	}

	err = t.ExecuteTemplate(w, "user.html", constants.WebData{
		Title:        constants.WebUserTitle,
		BaseURL:      constants.BaseURL,
		TemplateURL:  constants.TemplateUrl,
		LinkPageList: constants.LinkPageList,
	})
	if err != nil {
		fmt.Fprintf(w, "User Page Error, Please Try Again Later, Err : %v", err)
	}
}
