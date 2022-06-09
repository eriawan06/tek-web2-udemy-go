package entity

import "github.com/eriawan06/tek-web2-udemy-go/src/utils/common"

type Category struct {
	common.BaseEntity
	Name string `gorm:"type:varchar(255);uniqueIndex;not null"`
}
