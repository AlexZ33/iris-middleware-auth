package iris_middleware_auth

import "github.com/kataras/iris/v12"

var (
	PublicParties   []string
	PublicRoutes    []string
	ProtectedRoutes []string
	Global          []iris.Handler
)

func init() {
	PublicParties = []string{
		"/public",
		"/share",
	}
	PublicRoutes = []string{
		"/",
		"/login",
		"/auth/login",
		"/auth/grant",
		"/auth/access",
		"/auth/register",
	}
	ProtectedRoutes = []string{}
	Global = []iris.Handler{
		AuthRequired,
	}
}
