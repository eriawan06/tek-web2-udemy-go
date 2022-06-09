package service

import (
	ad "github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/category/model/dto"
)

type CategoryService interface {
	Create(claims ad.UserClaims, request dto.CreateCategoryRequest) error
	Update(claims ad.UserClaims, request dto.UpdateCategoryRequest, categoryId uint) error
	Delete(claims ad.UserClaims, categoryId uint) error
	GetAll() ([]dto.CategoryResponse, error)
	GetOne(categoryId uint) (dto.CategoryResponse, error)
}
