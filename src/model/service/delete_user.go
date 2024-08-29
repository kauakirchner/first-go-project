package service

import (
	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserServices(
	id string,
) *rest_err.RestErr {
	logger.Info("Init DeleteUser service", zap.String("journey", "deleteUser"))

	if err := ud.userRepository.DeleteUser(id); err != nil {
		logger.Error(
			"Error trying to call DeleteUser repository",
			err,
			zap.String("journey", "deleteUser"),
		)
		return err
	}

	logger.Info(
		"DeleteUser service executed succesfully",
		zap.String("userId", id),
		zap.String("journey", "deleteUser"),
	)
	return nil
}
