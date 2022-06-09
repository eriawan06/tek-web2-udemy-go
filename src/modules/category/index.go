package category

import (
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/category/controller"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/category/repository"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/category/service"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils/common"
	"gorm.io/gorm"
)

var (
	categoryController controller.CategoryController
)

type ModuleImplCategory struct {
	DB *gorm.DB
}

func New(database *gorm.DB) common.Module {
	return &ModuleImplCategory{DB: database}
}

func (module *ModuleImplCategory) InitModule() {
	categoryRepository := repository.NewCategoryRepository(module.DB)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController = controller.NewCategoryController(categoryService)
}

func GetController() controller.CategoryController {
	return categoryController
}
