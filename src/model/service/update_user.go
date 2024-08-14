package service

import (
	"github.com/kauakirchner/first-go-project/src/config/rest_err"
	"github.com/kauakirchner/first-go-project/src/model"
)

func (ud *userDomainService) UpdateUser(
	id string,
	userDomain model.InterfaceUserDomain,
) *rest_err.RestErr {
	return nil
}
