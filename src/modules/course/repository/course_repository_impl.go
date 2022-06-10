package repository

import (
	"errors"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course/model/entity"
	e "github.com/eriawan06/tek-web2-udemy-go/src/utils/errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type CourseRepositoryImpl struct {
	DB *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &CourseRepositoryImpl{DB: db}
}

func (repository *CourseRepositoryImpl) Create(course entity.Course) error {
	result := repository.DB.Create(&course)

	var mySqlErr *mysql.MySQLError
	if errors.As(result.Error, &mySqlErr) && mySqlErr.Number == 1062 {
		result.Error = e.ErrDuplicateKey
	}

	return result.Error
}

func (repository *CourseRepositoryImpl) Update(course entity.Course, courseId uint) error {
	result := repository.DB.Model(&entity.Course{}).Where("id=?", courseId).Updates(&course)
	return result.Error
}

func (repository *CourseRepositoryImpl) Delete(courseId uint) error {
	result := repository.DB.Delete(&entity.Course{}, courseId)
	return result.Error
}

func (repository *CourseRepositoryImpl) FindAll() ([]entity.CourseLite, error) {
	var courses []entity.CourseLite

	query := `
	SELECT c.id, c.code, c.name, c.excerpt,
		(select jsonb_agg(json_build_object(
				'id', cc.id,
				'category_id', cat.id,
				'category_name', cat.name))
			FROM course_categories cc
			LEFT JOIN categories cat ON cat.id = cc.category_id
			WHERE cc.course_code = c.code
		) categories
	FROM courses c
	`
	result := repository.DB.Raw(query).Scan(&courses)
	if result.Error != nil {
		return nil, result.Error
	}

	return courses, nil
}

func (repository *CourseRepositoryImpl) FindOne(courseId uint) (entity.Course, error) {
	var course entity.Course
	result := repository.DB.Where("id=?", courseId).First(&course)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result.Error = e.ErrDataNotFound
		}
		return course, result.Error
	}
	return course, nil
}

func (repository *CourseRepositoryImpl) FindOneDetail(courseId uint) (entity.CourseDetail, error) {
	var course entity.CourseDetail

	query := `
	SELECT c.*,
		(select jsonb_agg(json_build_object(
				'id', cc.id,
				'category_id', cat.id,
				'category_name', cat.name))
			FROM course_categories cc
			LEFT JOIN categories cat ON cat.id = cc.category_id
			WHERE cc.course_code = c.code
		) categories
	FROM courses c
	WHERE c.id = ?
	LIMIT 1
	`
	result := repository.DB.Raw(query, courseId).Scan(&course)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result.Error = e.ErrDataNotFound
		}
		return course, result.Error
	}

	return course, nil
}
