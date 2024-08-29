package service

import (
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
	"github.com/kauakirchner/first-go-project/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.InterfaceUserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.InterfaceUserRepository
}

type UserDomainService interface {
	CreateUserServices(model.InterfaceUserDomain) (model.InterfaceUserDomain, *rest_err.RestErr)
	UpdateUserServices(string, model.InterfaceUserDomain) *rest_err.RestErr
	DeleteUserServices(string) *rest_err.RestErr
	FindUserByIDServices(string) (model.InterfaceUserDomain, *rest_err.RestErr)
	FindUserByEmailServices(string) (model.InterfaceUserDomain, *rest_err.RestErr)
	LoginUserServices(model.InterfaceUserDomain) (model.InterfaceUserDomain, *rest_err.RestErr)
}
