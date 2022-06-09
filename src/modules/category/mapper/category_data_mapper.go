package mapper

import (
	ad "github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/category/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/category/model/entity"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils/common"
)

func CreateCategoryRequestToCategory(claims ad.UserClaims, request dto.CreateCategoryRequest) entity.Category {
	return entity.Category{
		Name: request.Name,
		BaseEntity: common.BaseEntity{
			CreatedBy: claims.Email,
			UpdatedBy: claims.Email,
		},
	}
}

func UpdateCategoryRequestToCategory(claims ad.UserClaims, request dto.UpdateCategoryRequest) entity.Category {
	return entity.Category{
		Name: request.Name,
		BaseEntity: common.BaseEntity{
			UpdatedBy: claims.Email,
		},
	}
}

func CategoryToCategoryResponse(category entity.Category) dto.CategoryResponse {
	return dto.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ListCategoryToListCategoryResponse(categories []entity.Category) []dto.CategoryResponse {
	var categoriesResp []dto.CategoryResponse
	for _, v := range categories {
		categoriesResp = append(categoriesResp, CategoryToCategoryResponse(v))
	}
	return categoriesResp
}
