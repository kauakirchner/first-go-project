package repository

import (
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(db *mongo.Database) InterfaceUserRepository {
	return &userRepository{
		db,
	}
}

type InterfaceUserRepository interface {
	CreteUser(
		userDomain model.InterfaceUserDomain,
	) (model.InterfaceUserDomain, *rest_err.RestErr)
}
