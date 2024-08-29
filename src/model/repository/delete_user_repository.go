package repository

import (
	"context"
	"os"

	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(
	id string,
) *rest_err.RestErr {
	logger.Info("Init DeleteUser repository", zap.String("journey", "deleteUser"))

	collection := ur.db.Collection(os.Getenv(DB_USER_COLLECTION))

	userIdHex, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	if _, err := collection.DeleteOne(context.Background(), filter); err != nil {
		logger.Error(
			"Error trying to delete user",
			err,
			zap.String("journey", "deleteUser"),
		)
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info(
		"DeleteUser repository executed succesfully",
		zap.String("userId", id),
		zap.String("journey", "deleteUser"),
	)
	return nil
}
