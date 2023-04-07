package response

import (
	"github.com/gin-gonic/gin"
)

var (
	SUCCESS    = 20000
	ERROR      = -1
	LOGINAGAIN = 20001 // 重新登录
)

func Success(c *gin.Context, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"code": SUCCESS,
		"msg":  msg,
		"data": data,
	})
}

func Error(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
