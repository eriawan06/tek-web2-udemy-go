package entity

import "github.com/eriawan06/tek-web2-udemy-go/src/utils/common"

type User struct {
	common.BaseEntity
	Name     string `gorm:"type:varchar(255);not null"`
	Email    string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Role     string `gorm:"type:varchar(20);not null"`
	//Participant *Participant `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
