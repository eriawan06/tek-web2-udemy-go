package service

import (
	ad "github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course/model/dto"
)

type CourseService interface {
	Create(claims ad.UserClaims, request dto.CreateCourseRequest) error
	Update(claims ad.UserClaims, request dto.UpdateCourseRequest, courseId uint) error
	Delete(claims ad.UserClaims, courseId uint) error
	GetAll() ([]dto.CourseResponse, error)
	GetOne(courseId uint) (dto.CourseDetailResponse, error)
}
