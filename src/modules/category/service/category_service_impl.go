package service

import (
	ad "github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/category/mapper"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/category/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/category/repository"
	e "github.com/eriawan06/tek-web2-udemy-go/src/utils/errors"
)

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{Repository: repository}
}

func (service CategoryServiceImpl) Create(claims ad.UserClaims, request dto.CreateCategoryRequest) error {
	if claims.Role != "admin" {
		return e.ErrForbidden
	}

	category := mapper.CreateCategoryRequestToCategory(claims, request)
	err := service.Repository.Create(category)
	return err
}

func (service CategoryServiceImpl) Update(claims ad.UserClaims, request dto.UpdateCategoryRequest, categoryId uint) error {
	if claims.Role != "admin" {
		return e.ErrForbidden
	}

	// check
	_, err := service.Repository.FindOne(categoryId)
	if err != nil {
		return err
	}

	category := mapper.UpdateCategoryRequestToCategory(claims, request)
	err = service.Repository.Update(category, categoryId)
	return err
}

func (service CategoryServiceImpl) Delete(claims ad.UserClaims, categoryId uint) error {
	if claims.Role != "admin" {
		return e.ErrForbidden
	}

	// check
	_, err := service.Repository.FindOne(categoryId)
	if err != nil {
		return err
	}

	err = service.Repository.Delete(categoryId)
	return err
}

func (service CategoryServiceImpl) GetAll() ([]dto.CategoryResponse, error) {
	categories, err := service.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	response := mapper.ListCategoryToListCategoryResponse(categories)
	return response, nil
}

func (service CategoryServiceImpl) GetOne(categoryId uint) (dto.CategoryResponse, error) {
	category, err := service.Repository.FindOne(categoryId)
	if err != nil {
		return dto.CategoryResponse{}, err
	}

	response := mapper.CategoryToCategoryResponse(category)
	return response, nil
}
