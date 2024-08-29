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

func (uc *interfaceUserController) LoginUser(ctx *gin.Context) {
	logger.Info("Init LoginUser controller", zap.String("journey", "loginUser"))

	var userRequest request.UserLogin

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error(
			"Error trying to validade user info",
			err,
			zap.String("journey", "loginUser"),
		)
		restErr := validation.ValidateUserError(err)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, err := uc.service.LoginUserServices(domain)

	if err != nil {
		logger.Error(
			"Error trying to call LoginUser service",
			err,
			zap.String("journey", "loginUser"),
		)
		ctx.JSON(err.Code, err)
		return
	}

	logger.Info(
		"LoginUser controller executed succesfully",
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser"),
	)
	ctx.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}
