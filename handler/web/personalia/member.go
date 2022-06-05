package personalia

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/golang_pos_app_service/entity/constants"
)

func WebMemberHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(constants.HeaderViewHtml, constants.LoadingViewHtml, constants.SideBarViewHtml, constants.FooterViewHtml, constants.MemberViewHtml)
	if err != nil {
		fmt.Fprintf(w, "Member Page Error, Please Try Again Later, Err : %v", err)
	}

	err = t.ExecuteTemplate(w, "member.html", constants.WebData{
		Title:        constants.WebMemberTitle,
		BaseURL:      constants.BaseURL,
		TemplateURL:  constants.TemplateUrl,
		LinkPageList: constants.LinkPageList,
	})
	if err != nil {
		fmt.Fprintf(w, "Member Page Error, Please Try Again Later, Err : %v", err)
	}
}
