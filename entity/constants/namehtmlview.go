package constants

const (
	HeaderViewHtml  = "./view/header.html"
	LoadingViewHtml = "./view/loading.html"
	SideBarViewHtml = "./view/sidebar.html"
	FooterViewHtml  = "./view/footer.html"

	HomeViewHtml  = "./view/home.html"
	UserViewHtml  = "./view/user.html"
	LoginViewHtml = "./view/login.html"
)

var LinkPageList = map[string]string{
	// Web Constanta
	"WebHome": BaseURL + "home",
	"WebUser": BaseURL + "user",

	// Services Constanta
	"ServiceGetAllUser": BaseURL + "svc/" + "users",
}
