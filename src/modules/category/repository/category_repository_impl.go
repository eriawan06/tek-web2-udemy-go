package repository

import (
	"errors"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/category/model/entity"
	e "github.com/eriawan06/tek-web2-udemy-go/src/utils/errors"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}

func (repository CategoryRepositoryImpl) Create(category entity.Category) error {
	result := repository.DB.Create(&category)
	var postgreErr *pgconn.PgError
	if errors.As(result.Error, &postgreErr) && postgreErr.Code == "23505" {
		result.Error = e.ErrDuplicateKey
	}
	return result.Error
}

func (repository CategoryRepositoryImpl) Update(category entity.Category, categoryId uint) error {
	result := repository.DB.Model(&entity.Category{}).Where("id=?", categoryId).Updates(&category)
	return result.Error
}

func (repository CategoryRepositoryImpl) Delete(categoryId uint) error {
	result := repository.DB.Delete(&entity.Category{}, categoryId)
	return result.Error
}

func (repository CategoryRepositoryImpl) FindAll() ([]entity.Category, error) {
	var categories []entity.Category

	result := repository.DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func (repository CategoryRepositoryImpl) FindOne(categoryId uint) (entity.Category, error) {
	var category entity.Category

	result := repository.DB.Where("id = ?", categoryId).First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result.Error = e.ErrDataNotFound
		}
		return category, result.Error
	}

	return category, nil
}
