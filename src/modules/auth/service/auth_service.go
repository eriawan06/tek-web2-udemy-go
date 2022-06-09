package service

import (
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/model/dto"
)

type AuthService interface {
	Register(request dto.RegisterRequest) error
	Login(request dto.LoginRequest) (dto.AuthResponse, error)
}
