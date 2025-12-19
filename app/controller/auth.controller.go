package controller

import (
	"gin-01/app/dto/request"
	"gin-01/app/helper"
	"gin-01/app/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(c *gin.Context)
}

type authController struct {
	authService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{authService: authService}
}

func (ctrl *authController) Register(c *gin.Context) {
	var req request.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ValidationErrorResponse(c, err.Error())
		return
	}

	data, err := ctrl.authService.Register(req)
	if err != nil {
		helper.ErrorResponse(c, 500, "failed to register", err.Error())
		return
	}

	helper.CreatedResponse(c, "user created successfully", data)
}
