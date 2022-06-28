package service

import (
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils"
)

type AuthService interface {
	Register(request dto.RegisterRequest) error
	Login(request dto.LoginRequest) (dto.AuthResponse, error)
	GoogleOauth(request utils.GoogleUserResult) (dto.AuthResponse, error)
}
