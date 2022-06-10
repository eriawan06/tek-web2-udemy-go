package repository

import (
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course/model/entity"
	"gorm.io/gorm"
)

type CourseCategoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCourseCategoryRepository(db *gorm.DB) CourseCategoryRepository {
	return &CourseCategoryRepositoryImpl{DB: db}
}

func (repository *CourseCategoryRepositoryImpl) Create(cc entity.CourseCategory) error {
	result := repository.DB.Create(&cc)
	return result.Error
}

func (repository *CourseCategoryRepositoryImpl) CreateBatch(ccs []entity.CourseCategory) error {
	result := repository.DB.Create(&ccs)
	return result.Error
}

func (repository *CourseCategoryRepositoryImpl) Delete(ccId uint) error {
	result := repository.DB.Delete(&entity.CourseCategory{}, ccId)
	return result.Error
}

func (repository *CourseCategoryRepositoryImpl) DeleteByCourseId(courseId uint) error {
	result := repository.DB.Where("course_id = ?", courseId).Delete(&entity.CourseCategory{})
	return result.Error
}
