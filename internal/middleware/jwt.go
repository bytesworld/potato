package middleware

import (
	"fmt"
	"gitee.com/bytesworld/tomato/configs"
	"gitee.com/bytesworld/tomato/internal/service"
	"gitee.com/bytesworld/tomato/pkg/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtAurh(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		fmt.Println(tokenStr)
		if tokenStr == "" {
			response.TokenFail(c)
			c.Abort()
			return
		}
		tokenStr = tokenStr[len(service.TokenType)+1:]
		fmt.Println(tokenStr)
		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &service.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(configs.AppObj.Config.Jwt.Secret), nil
		})
		fmt.Println(token)
		if err != nil {
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*service.CustomClaims)
		fmt.Println(claims)
		// Token 发布者校验
		if claims.Issuer != GuardName {
			response.TokenFail(c)
			c.Abort()
			return
		}

		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}
