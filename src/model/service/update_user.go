package service

import (
	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserServices(
	id string,
	userDomain model.InterfaceUserDomain,
) *rest_err.RestErr {
	logger.Info("Init UpdateUser service", zap.String("journey", "updateUser"))

	if err := ud.userRepository.UpdateUser(id, userDomain); err != nil {
		logger.Error(
			"Error trying to call repository",
			err,
			zap.String("journey", "updateUser"),
		)
		return err
	}

	logger.Info(
		"UpdateUser service executed succesfully",
		zap.String("userId", id),
		zap.String("journey", "updateUser"),
	)
	return nil
}
