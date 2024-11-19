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

// FindUserByID retrieves user information based on the provided user ID.
// @Summary Find User by ID
// @Description Retrieves user details based on the user ID provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userId path string true "ID of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: Invalid user ID"
// @Failure 404 {object} rest_err.RestErr "User not found"
// @Router /user/{userId} [get]
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

// FindUserByEmail retrieves user information based on the provided email.
// @Summary Find User by Email
// @Description Retrieves user details based on the email provided as a parameter.
// @Tags Users
// @Accept json
// @Produce json
// @Param userEmail path string true "Email of the user to be retrieved"
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {object} response.UserResponse "User information retrieved successfully"
// @Failure 400 {object} rest_err.RestErr "Error: Invalid user ID"
// @Failure 404 {object} rest_err.RestErr "User not found"
// @Router /getUserByEmail/{userEmail} [get]
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
