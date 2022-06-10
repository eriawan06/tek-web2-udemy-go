package database

import (
	cat "github.com/eriawan06/tek-web2-udemy-go/src/modules/category/model/entity"
	cou "github.com/eriawan06/tek-web2-udemy-go/src/modules/course/model/entity"
	ue "github.com/eriawan06/tek-web2-udemy-go/src/modules/user/model/entity"
	"gorm.io/gorm"
)

func MigrateDb(db *gorm.DB) {
	db.AutoMigrate(&ue.User{})
	db.AutoMigrate(&cat.Category{})
	db.AutoMigrate(&cou.Course{})
	db.AutoMigrate(&cou.CourseCategory{})
}
