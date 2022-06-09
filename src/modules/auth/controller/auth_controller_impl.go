package controller

import (
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/service"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils/common"
	e "github.com/eriawan06/tek-web2-udemy-go/src/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthControllerImpl struct {
	Service service.AuthService
}

func NewAuthController(service service.AuthService) AuthController {
	return &AuthControllerImpl{Service: service}
}

func (controller *AuthControllerImpl) Register(ctx *gin.Context) {
	var request dto.RegisterRequest

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

	//access register service
	err := controller.Service.Register(request)
	if err != nil {
		if err == e.ErrEmailAlreadyExists || err == e.ErrConfirmPasswordNotSame {
			common.SendError(ctx, http.StatusBadRequest, "Bad Request", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Register Success", nil)
}

func (controller *AuthControllerImpl) Login(ctx *gin.Context) {
	var request dto.LoginRequest

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

	// Access Services
	response, err := controller.Service.Login(request)
	if err != nil {
		if err == e.ErrWrongLoginCredential {
			common.SendError(ctx, http.StatusUnauthorized, "Unauthorized", []string{err.Error()})
			return
		}

		common.SendError(ctx, http.StatusInternalServerError, "Internal Server Error", []string{err.Error()})
		return
	}

	common.SendSuccess(ctx, http.StatusOK, "Login Success", &response)
}
