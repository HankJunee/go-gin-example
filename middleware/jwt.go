package middleware

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/util"
	"go-gin-example/util/response"
	"strings"
)

// JWTAuth 验证JWT
func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			response.Error(c, 20001, "", nil)
			c.Abort()
			return
		}

		jwt, err := util.ValidJWT(token)

		if err != nil {

			if strings.Contains(err.Error(), "token is expired") {
				// 登陆过期
				response.Error(c, 20001, "", nil)
			} else {
				// token不合法
				response.Error(c, 20001, "", nil)
			}

			c.Abort()
			return
		}

		c.Set("uid", int(jwt.Uid))

		c.Next()
	}
}
