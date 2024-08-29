package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/first-go-project/src/model"
	"github.com/kauakirchner/first-go-project/src/model/service"
)

var (
	InterfaceUserDomain model.InterfaceUserDomain
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
	FindUserByID(ctx *gin.Context)

	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}

type interfaceUserController struct {
	service service.UserDomainService
}
