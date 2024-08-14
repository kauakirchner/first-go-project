package service

import (
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.InterfaceUserDomain) *rest_err.RestErr
	UpdateUser(string, model.InterfaceUserDomain) *rest_err.RestErr
	FindUser(string) (*model.InterfaceUserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
