package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/validation"
	"github.com/kauakirchner/first-go-project/src/controller/model/request"
	"github.com/kauakirchner/first-go-project/src/controller/model/response"
	"go.uber.org/zap"
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
	logger.Info("User created succesfully", zap.String("journey", "createUser"))
	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}
	ctx.JSON(http.StatusOK, response)
}
