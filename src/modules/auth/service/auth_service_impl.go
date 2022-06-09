package service

import (
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/mapper"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/model/dto"
	ud "github.com/eriawan06/tek-web2-udemy-go/src/modules/user/model/dto"
	ue "github.com/eriawan06/tek-web2-udemy-go/src/modules/user/model/entity"
	ur "github.com/eriawan06/tek-web2-udemy-go/src/modules/user/repository"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils"
	e "github.com/eriawan06/tek-web2-udemy-go/src/utils/errors"
)

type AuthServiceImpl struct {
	UserRepo ur.UserRepository
}

func NewAuthService(userRepo ur.UserRepository) AuthService {
	return &AuthServiceImpl{UserRepo: userRepo}
}

func (service *AuthServiceImpl) Register(request dto.RegisterRequest) error {
	if request.Password != request.ConfirmPassword {
		return e.ErrConfirmPasswordNotSame
	}

	hashed, err := utils.HashPassword(request.Password)
	if err != nil {
		return err
	}
	request.Password = hashed

	// mapping
	user := mapper.RegisterRequestToUser(request)

	// Create new User
	err = service.UserRepo.Create(user)
	return err
}

func (service *AuthServiceImpl) Login(request dto.LoginRequest) (dto.AuthResponse, error) {
	var (
		user ue.User
		err  error
	)

	user, err = service.UserRepo.FindByEmail(request.Email)
	if err != nil {
		if err.Error() == "record not found" {
			return dto.AuthResponse{}, e.ErrWrongLoginCredential
		}
		return dto.AuthResponse{}, err
	}

	isPasswordValid := utils.CheckPasswordHash(request.Password, user.Password)
	if !isPasswordValid {
		return dto.AuthResponse{}, e.ErrWrongLoginCredential
	}

	// Generate Token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return dto.AuthResponse{}, err
	}

	// Return nil Error
	return dto.AuthResponse{
		Token: token,
		User: ud.UserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}
