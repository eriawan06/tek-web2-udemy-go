package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	ce "github.com/eriawan06/tek-web2-udemy-go/src/modules/category/model/entity"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils/common"
)

type CourseCategory struct {
	common.BaseEntity
	CourseCode string
	Course     Course `gorm:"foreignKey:CourseCode;references:Code;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CategoryID uint
	Category   ce.Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CourseCategoryDetail struct {
	Id           uint   `json:"id"`
	CategoryId   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
}

type CourseCategoryDetailList []CourseCategoryDetail

func (cc CourseCategoryDetailList) Value() (driver.Value, error) {
	return json.Marshal(cc)
}

func (cc *CourseCategoryDetailList) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	unmarshal := json.Unmarshal(b, &cc)
	return unmarshal
}
