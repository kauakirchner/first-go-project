package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var (
	DB_USER_COLLECTION = "DB_USER_COLLECTION"
)

type userRepository struct {
	db *mongo.Database
}

func (ur *userRepository) CreteUser(
	userDomain model.InterfaceUserDomain,
) (model.InterfaceUserDomain, *rest_err.RestErr) {
	logger.Info("Init create user repository", zap.String("journey", "userRepository"))

	collection := ur.db.Collection(os.Getenv(DB_USER_COLLECTION))

	value, err := userDomain.GetJSONValue()
	if err != nil {
		fmt.Println(err)
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))
	return userDomain, nil
}
