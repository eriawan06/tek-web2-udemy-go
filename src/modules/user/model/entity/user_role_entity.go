package entity

import "github.com/eriawan06/tek-web2-udemy-go/src/utils/common"

type UserRole struct {
	common.BaseEntity
	Name string `gorm:"type:varchar(20);uniqueIndex;not null"`
}
