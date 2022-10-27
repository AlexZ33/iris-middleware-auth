package iris_middleware_auth

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/dgrijalva/jwt-go"
)

func AuthRequired(ctx iris.Context) {
	path := ctx.Path()
	if ContainsPrefix(path, PublicParties) || StringArrayContains(PublicRoutes, path) {
		ctx.Next()
	} else if user, ok :=
}


