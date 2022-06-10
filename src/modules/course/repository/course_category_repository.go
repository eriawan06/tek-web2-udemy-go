package repository

import "github.com/eriawan06/tek-web2-udemy-go/src/modules/course/model/entity"

type CourseCategoryRepository interface {
	Create(cc entity.CourseCategory) error
	CreateBatch(ccs []entity.CourseCategory) error
	Delete(ccId uint) error
	DeleteByCourseId(courseId uint) error
}
