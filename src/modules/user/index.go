package user

import (
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/user/repository"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils/common"
	"gorm.io/gorm"
)

var (
	userRepository repository.UserRepository
)

type ModuleImplAuth struct {
	DB *gorm.DB
}

func New(database *gorm.DB) common.Module {
	return &ModuleImplAuth{DB: database}
}

func (module *ModuleImplAuth) InitModule() {
	userRepository = repository.NewUserRepository(module.DB)
}

func GetRepository() repository.UserRepository {
	return userRepository
}
