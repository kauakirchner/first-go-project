package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// DeleteUser deletes a user with the specified ID.
// @Summary Delete User
// @Description Deletes a user based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be deleted"
// @Success 200
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /user/{userId} [delete]
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
