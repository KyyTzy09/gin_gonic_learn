package service

import (
	"errors"
	"gin-01/app/dto/request"
	"gin-01/app/entity"
	"gin-01/app/helper"

	"gorm.io/gorm"
)

type AuthService interface {
	Register(req request.RegisterRequest) (any, error)
	Login(req request.LoginRequest) (string, error)
}

type authService struct {
	db *gorm.DB
}

func NewAuthService(db gorm.DB) AuthService {
	return &authService{db: &db}
}

func (s *authService) Register(req request.RegisterRequest) (any, error) {
	var existingUser *entity.UserEntity

	if err := s.db.First(&existingUser, "email = ?", req.Email).Error; err == nil {
		return nil, errors.New("email is already registered")
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := entity.UserEntity{
		Email:    req.Email,
		Password: &hashedPassword,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	return &user, nil
}

func (s *authService) Login(req request.LoginRequest) (string, error) {
	var existingUser *entity.UserEntity

	if err := s.db.First(&existingUser, "email = ?", req.Email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("user is not registered")
		}
		return "", err
	}

	if err := helper.ComparePassword(req.Password, *existingUser.Password); err != nil {
		return "", errors.New("incorrect password")
	}

	access_token, err := helper.GenerateToken(existingUser.UserId, existingUser.Email, "USER")
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return access_token, nil
}
