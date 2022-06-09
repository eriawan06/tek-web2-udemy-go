package dto

import (
	ud "github.com/eriawan06/tek-web2-udemy-go/src/modules/user/model/dto"
)

type RegisterRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserClaims struct {
	Authorized bool   `json:"authorized"`
	UserId     uint   `json:"user_id"`
	Email      string `json:"email"`
	RoleId     uint   `json:"role_id"`
	RoleName   string `json:"role_name"`
	Expired    int64  `json:"expired"`
}

type AuthResponse struct {
	Token string          `json:"token"`
	User  ud.UserResponse `json:"user"`
}
