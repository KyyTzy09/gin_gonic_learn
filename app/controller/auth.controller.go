package controller

import (
	"gin-01/app/dto/request"
	"gin-01/app/helper"
	"gin-01/app/service"

	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
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
func (ctrl *authController) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		helper.ValidationErrorResponse(c, err.Error())
		return
	}

	access_token, err := ctrl.authService.Login(req)
	if err != nil {
		helper.ErrorResponse(c, 500, "login failed", err.Error())
		return
	}

	helper.SuccessResponse(c, "login successfully", access_token)
}
