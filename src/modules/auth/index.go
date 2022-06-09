package auth

import (
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/controller"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/service"
	ur "github.com/eriawan06/tek-web2-udemy-go/src/modules/user/repository"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils/common"

	"gorm.io/gorm"
)

var (
	authService    service.AuthService
	authController controller.AuthController
)

type ModuleImplAuth struct {
	DB *gorm.DB
}

func New(database *gorm.DB) common.Module {
	return &ModuleImplAuth{DB: database}
}

func (module *ModuleImplAuth) InitModule() {
	userRepository := ur.NewUserRepository(module.DB)
	authService = service.NewAuthService(userRepository)
	authController = controller.NewAuthController(authService)
}

func GetController() controller.AuthController {
	return authController
}
