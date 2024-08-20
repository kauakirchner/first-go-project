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
	CreateUser(model.InterfaceUserDomain) (model.InterfaceUserDomain, *rest_err.RestErr)
	UpdateUser(string, model.InterfaceUserDomain) *rest_err.RestErr
	FindUser(string) (*model.InterfaceUserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
