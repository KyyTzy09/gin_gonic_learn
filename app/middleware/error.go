package middleware

import (
	"gin-01/app/helper"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic recovered: %v", err)
				helper.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err)
				c.Abort()
			}
		}()
		c.Next()
	}
}
