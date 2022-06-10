package mapper

import (
	ad "github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course/model/entity"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils/common"
)

func CreateCourseRequestToCourse(claims ad.UserClaims, request dto.CreateCourseRequest) entity.Course {
	return entity.Course{
		UserID:       claims.UserId,
		Code:         utils.GenerateUuid(),
		Name:         request.Name,
		Excerpt:      request.Excerpt,
		LearnSummary: request.LearnSummary,
		Requirement:  request.Requirement,
		Description:  request.Description,
		BaseEntity: common.BaseEntity{
			CreatedBy: claims.Email,
			UpdatedBy: claims.Email,
		},
	}
}

func UpdateCourseRequestToCourse(claims ad.UserClaims, request dto.UpdateCourseRequest) entity.Course {
	return entity.Course{
		Name:         request.Name,
		Excerpt:      request.Excerpt,
		LearnSummary: request.LearnSummary,
		Requirement:  request.Requirement,
		Description:  request.Description,
		BaseEntity: common.BaseEntity{
			UpdatedBy: claims.Email,
		},
	}
}

func CourseLiteToCourseResponse(cl entity.CourseLite) dto.CourseResponse {
	return dto.CourseResponse{
		Id:         cl.Id,
		Code:       cl.Code,
		Name:       cl.Name,
		Excerpt:    cl.Excerpt,
		Categories: cl.Categories,
	}
}

func ListCourseLiteToListCourseResponse(cls []entity.CourseLite) []dto.CourseResponse {
	var coursesResp []dto.CourseResponse
	for _, course := range cls {
		coursesResp = append(coursesResp, CourseLiteToCourseResponse(course))
	}
	return coursesResp
}

func CourseDetailToCourseDetailResponse(cd entity.CourseDetail) dto.CourseDetailResponse {
	return dto.CourseDetailResponse{
		Id:           cd.Id,
		Code:         cd.Code,
		Name:         cd.Name,
		Excerpt:      cd.Excerpt,
		LearnSummary: cd.LearnSummary,
		Requirement:  cd.Requirement,
		Description:  cd.Description,
		Categories:   cd.Categories,
	}
}

func CourseCategoryDetailToCourseCategoryDetailResponse(cc entity.CourseCategoryDetail) dto.CourseCategoryDetailResponse {
	return dto.CourseCategoryDetailResponse{
		CategoryId:   cc.CategoryId,
		CategoryName: cc.CategoryName,
	}
}

func ListCourseCategoryDetailToListCourseCategoryDetailResponse(ccs []entity.CourseCategoryDetail) []dto.CourseCategoryDetailResponse {
	var ccsResp []dto.CourseCategoryDetailResponse
	for _, cc := range ccs {
		ccsResp = append(ccsResp, CourseCategoryDetailToCourseCategoryDetailResponse(cc))
	}
	return ccsResp
}

func BuildCourseCategories(courseCode string, categoriesId []uint) []entity.CourseCategory {
	var courseCategories []entity.CourseCategory

	for _, categoryId := range categoriesId {
		courseCategories = append(courseCategories, entity.CourseCategory{
			CourseCode: courseCode,
			CategoryID: categoryId,
		})
	}

	return courseCategories
}
