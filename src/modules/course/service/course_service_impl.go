package service

import (
	ad "github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course/mapper"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course/repository"
	e "github.com/eriawan06/tek-web2-udemy-go/src/utils/errors"
)

type CourseServiceImpl struct {
	Repository         repository.CourseRepository
	CourseCategoryRepo repository.CourseCategoryRepository
}

func NewCourseService(
	repository repository.CourseRepository,
	courseCategoryRepository repository.CourseCategoryRepository) CourseService {
	return &CourseServiceImpl{
		Repository:         repository,
		CourseCategoryRepo: courseCategoryRepository,
	}
}

func (service CourseServiceImpl) Create(claims ad.UserClaims, request dto.CreateCourseRequest) error {
	if claims.Role != "user" {
		return e.ErrForbidden
	}

	//TODO: DB Transaction

	//create course
	course := mapper.CreateCourseRequestToCourse(claims, request)
	err := service.Repository.Create(course)
	if err != nil {
		return err
	}

	//create course categories
	courseCategories := mapper.BuildCourseCategories(course.Code, request.Categories)
	err = service.CourseCategoryRepo.CreateBatch(courseCategories)
	if err != nil {
		return err
	}

	return nil
}

func (service CourseServiceImpl) Update(claims ad.UserClaims, request dto.UpdateCourseRequest, courseId uint) error {
	if claims.Role != "user" {
		return e.ErrForbidden
	}

	// check course
	checkCourse, err := service.Repository.FindOne(courseId)
	if err != nil {
		return err
	}

	// check course's owner
	if checkCourse.UserID != claims.UserId {
		return e.ErrNotTheOwner
	}

	// update course
	course := mapper.UpdateCourseRequestToCourse(claims, request)
	err = service.Repository.Update(course, courseId)

	return err
}

func (service CourseServiceImpl) Delete(claims ad.UserClaims, courseId uint) error {
	if claims.Role != "user" {
		return e.ErrForbidden
	}

	// check course
	checkCourse, err := service.Repository.FindOne(courseId)
	if err != nil {
		return err
	}

	// check course's owner
	if checkCourse.UserID != claims.UserId {
		return e.ErrNotTheOwner
	}

	// delete course
	err = service.Repository.Delete(courseId)
	return err
}

func (service CourseServiceImpl) GetAll() ([]dto.CourseResponse, error) {
	courses, err := service.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	coursesResponse := mapper.ListCourseLiteToListCourseResponse(courses)
	return coursesResponse, err
}

func (service CourseServiceImpl) GetOne(courseId uint) (dto.CourseDetailResponse, error) {
	course, err := service.Repository.FindOneDetail(courseId)
	if err != nil {
		return dto.CourseDetailResponse{}, err
	}

	courseResponse := mapper.CourseDetailToCourseDetailResponse(course)
	return courseResponse, err
}
