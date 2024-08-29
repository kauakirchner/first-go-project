package service

import (
	"github.com/kauakirchner/first-go-project/src/config/logger"
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (model.InterfaceUserDomain, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail service", zap.String("journey", "findUserByEmail"))

	return ud.userRepository.FindUserByEmail(email)
}

func (ud *userDomainService) FindUserByIDServices(
	id string,
) (model.InterfaceUserDomain, *rest_err.RestErr) {
	logger.Info("Init findUserByID service", zap.String("journey", "findUserByID"))

	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) findUserByEmailAndPasswordServices(
	email,
	password string,
) (model.InterfaceUserDomain, *rest_err.RestErr) {
	logger.Info("Init findUserByID service", zap.String("journey", "findUserByEmailAndPassword"))

	return ud.userRepository.FindUserByEmailAndPassword(email, password)
}
