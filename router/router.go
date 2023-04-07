package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/api"
	"go-gin-example/middleware"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.Cors())

	return router
}

func routers(r *gin.Engine) {
	publicRouter := r.Group("")
	{
		publicRouter.POST("register", api.Register)
	}

	privateRouter := r.Group("").Use(middleware.JWTAuth())
	{
		privateRouter.GET("home", api.Home)
	}
}
