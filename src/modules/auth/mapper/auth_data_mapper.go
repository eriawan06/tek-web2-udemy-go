package mapper

import (
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/auth/model/dto"
	"github.com/eriawan06/tek-web2-udemy-go/src/modules/user/model/entity"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils/common"
	"github.com/eriawan06/tek-web2-udemy-go/src/utils/constants"
)

func RegisterRequestToUser(request dto.RegisterRequest) entity.User {
	return entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Role:     request.Role,
		BaseEntity: common.BaseEntity{
			CreatedBy: "self",
			UpdatedBy: "self",
		},
	}
}

func GoogleUserResultToUser(request utils.GoogleUserResult) entity.User {
	return entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Role:     "user",
		AuthType: constants.AuthTypeGoogle,
		BaseEntity: common.BaseEntity{
			CreatedBy: "self",
			UpdatedBy: "self",
		},
	}
}
