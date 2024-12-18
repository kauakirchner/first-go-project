package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/validation"
	"github.com/kauakirchner/first-go-project/src/controller/model/request"
	"github.com/kauakirchner/first-go-project/src/model"
	"github.com/kauakirchner/first-go-project/src/view"
	"go.uber.org/zap"
)

// CreateUser Creates a new user
// @Summary Create a new user
// @Description Create a new user with the provided user information
// @Tags Users
// @Accept json
// @Produce json
// @Param userRequest body request.UserRequest true "User information for registration"
// @Success 200 {object} response.UserResponse
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /user [post]
func (uc *interfaceUserController) CreateUser(ctx *gin.Context) {
	logger.Info("Init CreateUser controller", zap.String("journey", "createUser"))
	var userRequest request.UserRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error(
			"Error trying to validade user info",
			err,
			zap.String("journey", "createUser"),
		)
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

	domainResult, err := uc.service.CreateUserServices(domain)

	if err != nil {
		logger.Error(
			"Error trying to call CreateUser service",
			err,
			zap.String("journey", "createUser"),
		)
		ctx.JSON(err.Code, err)
		return
	}

	logger.Info(
		"CreateUser controller executed succesfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "createUser"),
	)
	ctx.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}
