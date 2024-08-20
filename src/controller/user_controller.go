package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/first-go-project/src/model/service"
)

func NewUserControllerInterface(
	interfaceService service.UserDomainService,
) InterfaceUserController {
	return &interfaceUserController{
		service: interfaceService,
	}
}

type InterfaceUserController interface {
	FindUserByEmail(ctx *gin.Context)
	FindUserById(ctx *gin.Context)

	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

type interfaceUserController struct {
	service service.UserDomainService
}
