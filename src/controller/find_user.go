package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *interfaceUserController) FindUserByID(ctx *gin.Context) {
	logger.Info("Init findUserByID controller", zap.String("journey", "findUserByID"))

	userId := ctx.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error(
			"Error trying to validade userID",
			err,
			zap.String("journey", "findUserByID"),
		)
		errorMessage := rest_err.NewBadRequestError(
			"UserID is not a valid ID",
		)
		ctx.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)

	if err != nil {
		logger.Error(
			"Error trying to call findUserById service",
			err,
			zap.String("journey", "findUserByID"),
		)
		ctx.JSON(err.Code, err)
	}

	logger.Info(
		"findUserById controller executed succesfully",
		zap.String("journey", "findUserByID"),
	)
	ctx.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}

func (uc *interfaceUserController) FindUserByEmail(ctx *gin.Context) {
	logger.Info("Init findUserByEmail controller", zap.String("journey", "findUserByEmail"))

	userEmail := ctx.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error(
			"Error trying to validade userEmail",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		errorMessage := rest_err.NewBadRequestError(
			"UserEmail is not a valid email",
		)
		ctx.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)

	if err != nil {
		logger.Error(
			"Error trying to call findUserByEmail service",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		ctx.JSON(err.Code, err)
	}

	logger.Info(
		"findUserByEmail controller executed succesfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", userEmail),
	)

	ctx.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}
