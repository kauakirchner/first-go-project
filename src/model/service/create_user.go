package service

import (
	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(
	userDomain model.InterfaceUserDomain,
) (model.InterfaceUserDomain, *rest_err.RestErr) {
	logger.Info("Init CreateUser service", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error(
			"Error trying to call repository",
			err,
			zap.String("journey", "createUser"),
		)
		return nil, err
	}

	logger.Info(
		"CreateUser service executed succesfully",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"),
	)
	return userDomainRepository, nil
}
