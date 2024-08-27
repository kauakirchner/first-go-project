package repository

import (
	"context"
	"os"

	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
	"github.com/kauakirchner/first-go-project/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(
	userDomain model.InterfaceUserDomain,
) (model.InterfaceUserDomain, *rest_err.RestErr) {
	logger.Info("Init CreateUser repository", zap.String("journey", "createUser"))

	collection := ur.db.Collection(os.Getenv(DB_USER_COLLECTION))

	value := converter.ConvertDomainToEntity(userDomain)
	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		logger.Error(
			"Error trying to create user",
			err,
			zap.String("journey", "createUser"),
		)
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	value.ID = result.InsertedID.(primitive.ObjectID)
	logger.Info(
		"CreateUser repository executed succesfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"),
	)
	return converter.ConvertEntityToDomain(*value), nil
}
