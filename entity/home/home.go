package home

const WebHomeTitle = "Home"

type WebHomeData struct {
	Title       string `default:"Home"`
	BaseURL     string
	TemplateURL string
}
