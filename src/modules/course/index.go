package course

import (
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course/controller"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course/repository"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course/service"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils/common"
	"gorm.io/gorm"
)

var (
	courseController controller.CourseController
)

type ModuleImplCourse struct {
	DB *gorm.DB
}

func New(database *gorm.DB) common.Module {
	return &ModuleImplCourse{DB: database}
}

func (module *ModuleImplCourse) InitModule() {
	courseCategoryRepository := repository.NewCourseCategoryRepository(module.DB)
	courseRepository := repository.NewCourseRepository(module.DB)
	courseService := service.NewCourseService(courseRepository, courseCategoryRepository)
	courseController = controller.NewCourseController(courseService)
}

func GetController() controller.CourseController {
	return courseController
}
