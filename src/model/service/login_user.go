package service

import (
	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(
	userDomain model.InterfaceUserDomain,
) (model.InterfaceUserDomain, string, *rest_err.RestErr) {
	logger.Info("Init LoginUser service", zap.String("journey", "loginUser"))

	userDomain.EncryptPassword()
	user, err := ud.findUserByEmailAndPasswordServices(
		userDomain.GetEmail(),
		userDomain.GetPassword(),
	)

	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()

	if err != nil {
		return nil, "", err
	}

	logger.Info(
		"LoginUser service executed succesfully",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"),
	)
	return user, token, nil
}
