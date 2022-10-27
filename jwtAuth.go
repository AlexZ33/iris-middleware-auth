package iris_middleware_auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
)

// ParseAdminToken 解析admin token
// 基于jwt
func ParseAdminToken(ctx iris.Context, TokenSecret string) {
	tokenString := ctx.GetHeader("admin")
	token, tokenErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			// can not parse the token
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(TokenSecret), nil
	})
}
