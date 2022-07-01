package entity

import (
	ue "github.com/eriawan06/tek-web2-udemy-go/src/modules/user/model/entity"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils/common"
)

type Course struct {
	common.BaseEntity
	UserID       uint
	User         ue.User
	Code         string  `gorm:"type:varchar(36);uniqueIndex;not null"`
	Name         string  `gorm:"type:varchar(60);not null"`
	Excerpt      string  `gorm:"type:varchar(255);not null"`
	LearnSummary string  `gorm:"type:text;not null"`
	Requirement  *string `gorm:"type:text"`
	Description  *string `gorm:"type:text"`
	CoverImage   *string `gorm:"type:varchar(255);null"`
}

type CourseLite struct {
	Id         uint
	Code       string
	Name       string
	Excerpt    string
	CoverImage *string
	Author     string
	Categories CourseCategoryDetailList
}

type CourseDetail struct {
	Id           uint
	UserID       uint
	Code         string
	Name         string
	Excerpt      string
	LearnSummary string
	Requirement  *string
	Description  *string
	CoverImage   *string
	Categories   CourseCategoryDetailList
	//Modules
	//Reviews
}
