package router

import (
	"gin-01/app/controller"
	"gin-01/app/middleware"
	"gin-01/app/service"
	"gin-01/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if config.AppConfig.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.ErrorHandler())

	authService := service.NewAuthService(*config.GetDb())

	authController := controller.NewAuthController(authService)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("auth")
		{
			auth.POST("register", authController.Register)
		}
	}
	return r
}
