package repository

import "github.com/eriawan06/tek-web2-udemy-go/src/modules/user/model/entity"

type UserRepository interface {
	Create(user entity.User) error
	Update(user entity.User, userId uint) error
	Delete(userId uint) error
	FindAll() ([]entity.User, error)
	FindById(userId uint) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}
