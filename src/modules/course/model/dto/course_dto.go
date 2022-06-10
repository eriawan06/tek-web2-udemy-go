package dto

import "github.com/eriawan06/tek-web2-udemy-go/src/modules/course/model/entity"

type CreateCourseRequest struct {
	Name         string  `json:"name" binding:"required"`
	Excerpt      string  `json:"excerpt" binding:"required"`
	LearnSummary string  `json:"learn_summary" binding:"required"`
	Requirement  *string `json:"requirement"`
	Description  *string `json:"description"`
	Categories   []uint  `json:"categories"` //slice of category_ids
}

type UpdateCourseRequest struct {
	Name         string  `json:"name" binding:"required"`
	Excerpt      string  `json:"excerpt" binding:"required"`
	LearnSummary string  `json:"learn_summary" binding:"required"`
	Requirement  *string `json:"requirement"`
	Description  *string `json:"description"`
}

type CourseResponse struct {
	Id         uint                            `json:"id"`
	Code       string                          `json:"course_code"`
	Name       string                          `json:"name"`
	Excerpt    string                          `json:"excerpt"`
	Categories entity.CourseCategoryDetailList `json:"categories"`
}

type CourseDetailResponse struct {
	Id           uint                            `json:"id"`
	Code         string                          `json:"course_code"`
	UserId       uint                            `json:"user_id"`
	Name         string                          `json:"name"`
	Excerpt      string                          `json:"excerpt"`
	LearnSummary string                          `json:"learn_summary"`
	Requirement  *string                         `json:"requirement"`
	Description  *string                         `json:"description"`
	Categories   entity.CourseCategoryDetailList `json:"categories"`
}
