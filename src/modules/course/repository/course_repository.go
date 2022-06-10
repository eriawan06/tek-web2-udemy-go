package repository

import "github.com/eriawan06/tek-web2-udemy-go/src/modules/course/model/entity"

type CourseRepository interface {
	Create(course entity.Course) error
	Update(course entity.Course, courseId uint) error
	Delete(courseId uint) error
	FindAll() ([]entity.CourseLite, error)
	FindOne(courseId uint) (entity.Course, error)
	FindOneDetail(courseId uint) (entity.CourseDetail, error)
}
