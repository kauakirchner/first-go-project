package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/config/validation"
	"github.com/kauakirchner/first-go-project/src/controller/model/request"
	"github.com/kauakirchner/first-go-project/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// UpdateUser updates user information with the specified ID.
// @Summary Update User
// @Description Updates user details based on the ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be updated"
// @Param userRequest body request.UserUpdateRequest true "User information for update"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200
// @Failure 400 {object} rest_err.RestErr
// @Failure 500 {object} rest_err.RestErr
// @Router /user/{userId} [put]
func (uc *interfaceUserController) UpdateUser(ctx *gin.Context) {
	logger.Info("Init UpdateUser controller", zap.String("journey", "updateUser"))
	var userRequest request.UserUpdateRequest

	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		logger.Error(
			"Error trying to validade user info",
			err,
			zap.String("journey", "updateUser"),
		)
		restErr := validation.ValidateUserError(err)
		ctx.JSON(restErr.Code, restErr)
		return
	}

	userId := ctx.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error(
			"Error trying to validade userId",
			err,
			zap.String("journey", "updateUser"),
		)
		errRest := rest_err.NewBadRequestError("Invalid userId, must be a hex value")
		ctx.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	if err := uc.service.UpdateUserServices(userId, domain); err != nil {
		logger.Error(
			"Error trying to call UpdateUser service",
			err,
			zap.String("journey", "updateUser"),
		)
		ctx.JSON(err.Code, err)
		return
	}

	logger.Info(
		"UpdateUser controller executed succesfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
	)
	ctx.Status(http.StatusOK)
}
