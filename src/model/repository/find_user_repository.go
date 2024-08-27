package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
	"github.com/kauakirchner/first-go-project/src/model/repository/entity"
	"github.com/kauakirchner/first-go-project/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.InterfaceUserDomain, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail user repository", zap.String("journey", "findUserByEmail"))

	collection := ur.db.Collection(os.Getenv(DB_USER_COLLECTION))

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(
				errorMessage,
				err,
				zap.String("journey", "findUserByEmail"),
			)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by email"
		logger.Error(
			errorMessage,
			err,
			zap.String("journey", "findUserByEmail"),
		)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info(
		"findUserByEmail repository executed succesfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("email", email),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.InterfaceUserDomain, *rest_err.RestErr) {
	logger.Info("Init findUserByID user repository", zap.String("journey", "findUserByID"))

	collection := ur.db.Collection(os.Getenv(DB_USER_COLLECTION))

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objectId}}

	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this id: %s", id)
			logger.Error(
				errorMessage,
				err,
				zap.String("journey", "findUserByID"),
			)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Error trying to find user by id"
		logger.Error(
			errorMessage,
			err,
			zap.String("journey", "findUserByID"),
		)
		return nil, rest_err.NewInternalServerError(errorMessage)
	}
	logger.Info(
		"findUserByID repository executed succesfully",
		zap.String("journey", "findUserByID"),
		zap.String("userId", userEntity.ID.Hex()),
	)
	return converter.ConvertEntityToDomain(*userEntity), nil
}
