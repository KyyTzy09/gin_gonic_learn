package router

import (
	"gin-01/config"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	if config.AppConfig.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Server is running",
		})
	})

	return r
}
