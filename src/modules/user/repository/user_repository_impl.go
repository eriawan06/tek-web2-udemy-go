package repository

import (
	"errors"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/user/model/entity"
	e "github.com/eriawan06/tek-web2-udemy-go/src/utils/errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (repository *UserRepositoryImpl) Create(user entity.User) error {
	result := repository.DB.Create(&user)

	var mySqlErr *mysql.MySQLError
	if errors.As(result.Error, &mySqlErr) && mySqlErr.Number == 1062 {
		if strings.Contains(mySqlErr.Message, "idx_users_email") {
			result.Error = e.ErrEmailAlreadyExists
		}
	}

	return result.Error
}

func (repository *UserRepositoryImpl) Update(user entity.User, userId uint) error {
	result := repository.DB.Model(&entity.User{}).Where("id = ?", userId).Updates(&user)
	return result.Error
}

func (repository *UserRepositoryImpl) Delete(userId uint) error {
	result := repository.DB.Delete(&entity.User{}, userId)
	return result.Error
}

func (repository *UserRepositoryImpl) FindAll() ([]entity.User, error) {
	var users []entity.User
	result := repository.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (repository *UserRepositoryImpl) FindById(userId uint) (entity.User, error) {
	var user entity.User

	result := repository.DB.Where("id = ?", userId).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result.Error = e.ErrDataNotFound
		}
		return user, result.Error
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	result := repository.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			result.Error = e.ErrEmailNotRegistered
		}
		return user, result.Error
	}

	return user, nil
}
