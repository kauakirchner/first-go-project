package service

import (
	"fmt"

	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.InterfaceUserDomain,
) *rest_err.RestErr {
	logger.Info("Init create user model", zap.String("journey", "createUser"))
	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}
