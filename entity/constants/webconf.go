package constants

const WebHomeTitle = "Home"
const WebUserTitle = "User"
const WebMemberTitle = "Member"

type WebData struct {
	Title        string
	BaseURL      string
	TemplateURL  string
	LinkPageList map[string]string
}
