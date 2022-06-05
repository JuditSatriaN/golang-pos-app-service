package constants

const (
	HeaderViewHtml  = "./view/header.html"
	LoadingViewHtml = "./view/loading.html"
	SideBarViewHtml = "./view/sidebar.html"
	FooterViewHtml  = "./view/footer.html"

	HomeViewHtml   = "./view/home.html"
	LoginViewHtml  = "./view/login.html"
	UserViewHtml   = "./view/user.html"
	MemberViewHtml = "./view/member.html"
)

var LinkPageList = map[string]string{
	// Web Constanta
	"WebHome":   BaseURL + "home",
	"WebUser":   BaseURL + "user",
	"WebMember": BaseURL + "member",

	// Services Constanta
	"ServiceGetAllUser": BaseURL + "svc/" + "users",
}
