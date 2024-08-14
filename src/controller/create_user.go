package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/validation"
	"github.com/kauakirchner/first-go-project/src/controller/model/request"
	"github.com/kauakirchner/first-go-project/src/model"
	"github.com/kauakirchner/first-go-project/src/model/service"
	"go.uber.org/zap"
)

var (
	InterfaceUserDomain model.InterfaceUserDomain
)

func CreateUser(ctx *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validade user info", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		ctx.JSON(restErr.Code, restErr)
		return
	}
	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	service := service.NewUserDomainService()

	if err := service.CreateUser(domain); err != nil {
		ctx.JSON(err.Code, err)
		return
	}

	logger.Info("User created succesfully", zap.String("journey", "createUser"))
	ctx.String(http.StatusOK, "empty")
}
