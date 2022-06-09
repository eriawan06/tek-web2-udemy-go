package repository

import "github.com/eriawan06/tek-web2-udemy-go/src/modules/category/model/entity"

type CategoryRepository interface {
	Create(category entity.Category) error
	Update(category entity.Category, categoryId uint) error
	Delete(categoryId uint) error
	FindAll() ([]entity.Category, error)
	FindOne(categoryId uint) (entity.Category, error)
}
