package repository

import (
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db *mongo.Database
}

const (
	DB_USER_COLLECTION = "DB_USER_COLLECTION"
)

func NewUserRepository(db *mongo.Database) InterfaceUserRepository {
	return &userRepository{
		db,
	}
}

type InterfaceUserRepository interface {
	CreateUser(
		userDomain model.InterfaceUserDomain,
	) (model.InterfaceUserDomain, *rest_err.RestErr)

	FindUserByID(
		id string,
	) (model.InterfaceUserDomain, *rest_err.RestErr)

	FindUserByEmail(
		id string,
	) (model.InterfaceUserDomain, *rest_err.RestErr)

	UpdateUser(
		id string,
		userDomain model.InterfaceUserDomain,
	) *rest_err.RestErr
}
