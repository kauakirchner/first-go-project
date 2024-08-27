package repository

import (
	"context"
	"os"

	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
	"github.com/kauakirchner/first-go-project/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUser(
	id string,
	userDomain model.InterfaceUserDomain,
) *rest_err.RestErr {
	logger.Info("Init UpdateUser repository", zap.String("journey", "updateUser"))

	collection := ur.db.Collection(os.Getenv(DB_USER_COLLECTION))

	value := converter.ConvertDomainToEntity(userDomain)

	userIdHex, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: userIdHex}}
	update := bson.D{{Key: "$set", Value: value}}

	if _, err := collection.UpdateOne(context.Background(), filter, update); err != nil {
		logger.Error(
			"Error trying to update user",
			err,
			zap.String("journey", "updateUser"),
		)
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"UpdateUser repository executed succesfully",
		zap.String("userId", id),
		zap.String("journey", "updateUser"),
	)
	return nil
}
