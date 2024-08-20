package view

import (
	"github.com/kauakirchner/first-go-project/src/controller/model/response"
	"github.com/kauakirchner/first-go-project/src/model"
)

func ConvertDomainToResponse(
	userDomain model.InterfaceUserDomain,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
