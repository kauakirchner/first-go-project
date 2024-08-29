package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *interfaceUserController) DeleteUser(ctx *gin.Context) {
	logger.Info("Init DeleteUser controller", zap.String("journey", "deleteUser"))

	userId := ctx.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error(
			"Error trying to validade userId",
			err,
			zap.String("journey", "deleteUser"),
		)
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		ctx.JSON(errRest.Code, errRest)
	}

	if err := uc.service.DeleteUserServices(userId); err != nil {
		logger.Error(
			"Error trying to call DeleteUser service",
			err,
			zap.String("journey", "deleteUser"),
		)
		ctx.JSON(err.Code, err)
		return
	}

	logger.Info(
		"DeleteUser controller executed succesfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)
	ctx.Status(http.StatusOK)
}
