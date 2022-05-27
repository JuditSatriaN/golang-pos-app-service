package constants

const WebHomeTitle = "Home"
const WebUserTitle = "User"

type WebData struct {
	Title        string
	BaseURL      string
	TemplateURL  string
	LinkPageList map[string]string
}
