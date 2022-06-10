package controller

import (
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/course/service"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils/common"
	e "github.com/eriawan06/tek-web2-udemy-go/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CourseControllerImpl struct {
	Service service.CourseService
}

func NewCourseController(service service.CourseService) CourseController {
	return &CourseControllerImpl{Service: service}
}

func (controller CourseControllerImpl) Create(ctx *gin.Context) {
	userClaims, err := utils.GetUserCredentialFromToken(ctx)
	if err != nil {
		common.SendError(ctx, http.StatusUnauthorized, "Invalid Token", []string{err.Error()})
		return
	}

	var request dto.CreateCourseRequest
	errorBinding := ctx.ShouldBindJSON(&request)
	if errorBinding != nil {
		// Check if there is EOF error
		if errorBinding.Error() == "EOF" {
			common.SendError(ctx, http.StatusBadRequest, "Body is empty", []string{"Body required"})
			return
		}

		// When Binding Error
		common.SendError(ctx, http.StatusBadRequest, "Invalid request", utils.SplitError(errorBinding))
		return
	}

	err = controller.Service.Create(userClaims, request)
	if err != nil {
		if err == e.ErrForbidden {
			common.SendError(ctx, http.StatusForbidden, "Forbidden", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusCreated, "Create Course Success", nil)
}

func (controller CourseControllerImpl) Update(ctx *gin.Context) {
	userClaims, err := utils.GetUserCredentialFromToken(ctx)
	if err != nil {
		common.SendError(ctx, http.StatusUnauthorized, "Invalid Token", []string{err.Error()})
		return
	}

	var request dto.UpdateCourseRequest
	errorBinding := ctx.ShouldBindJSON(&request)
	if errorBinding != nil {
		// Check if there is EOF error
		if errorBinding.Error() == "EOF" {
			common.SendError(ctx, http.StatusBadRequest, "Body is empty", []string{"Body required"})
			return
		}

		// When Binding Error
		common.SendError(ctx, http.StatusBadRequest, "Invalid request", utils.SplitError(errorBinding))
		return
	}

	courseId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	err = controller.Service.Update(userClaims, request, uint(courseId))
	if err != nil {
		if err == e.ErrDataNotFound {
			common.SendError(ctx, http.StatusNotFound, "Data Not Found", []string{err.Error()})
			return
		}

		if err == e.ErrForbidden || err == e.ErrNotTheOwner {
			common.SendError(ctx, http.StatusForbidden, "Forbidden", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Update Course Success", nil)
}

func (controller CourseControllerImpl) Delete(ctx *gin.Context) {
	userClaims, err := utils.GetUserCredentialFromToken(ctx)
	if err != nil {
		common.SendError(ctx, http.StatusUnauthorized, "Invalid Token", []string{err.Error()})
		return
	}

	courseId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	err = controller.Service.Delete(userClaims, uint(courseId))
	if err != nil {
		if err == e.ErrDataNotFound {
			common.SendError(ctx, http.StatusNotFound, "Data Not Found", []string{err.Error()})
			return
		}

		if err == e.ErrForbidden || err == e.ErrNotTheOwner {
			common.SendError(ctx, http.StatusForbidden, "Forbidden", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Delete Course Success", nil)
}

func (controller CourseControllerImpl) GetAll(ctx *gin.Context) {
	data, err := controller.Service.GetAll()
	if err != nil {
		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Get All Course Success", data)
}

func (controller CourseControllerImpl) GetOne(ctx *gin.Context) {
	courseId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		common.SendError(ctx, http.StatusBadRequest, "Invalid Id", []string{err.Error()})
		return
	}

	data, err := controller.Service.GetOne(uint(courseId))
	if err != nil {
		if err == e.ErrDataNotFound {
			common.SendError(ctx, http.StatusNotFound, "Data Not Found", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Get One Course Success", data)
}
