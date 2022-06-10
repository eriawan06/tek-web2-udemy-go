package dto

type CourseCategoryDetailResponse struct {
	Id           uint   `json:"course_category_id"`
	CategoryId   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
}
