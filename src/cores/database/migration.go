package database

import (
	ce "github.com/eriawan06/tek-web2-udemy-go/src/modules/category/model/entity"
	ue "github.com/eriawan06/tek-web2-udemy-go/src/modules/user/model/entity"
	"gorm.io/gorm"
)

func MigrateDb(db *gorm.DB) {
	db.AutoMigrate(&ue.User{})
	db.AutoMigrate(&ce.Category{})
}
