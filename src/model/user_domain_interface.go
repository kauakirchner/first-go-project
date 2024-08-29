package model

import "github.com/kauakirchner/first-go-project/src/config/rest_err"

type InterfaceUserDomain interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetID() string

	SetID(string)

	EncryptPassword()
	GenerateToken() (string, *rest_err.RestErr)
}

func NewUserDomain(
	email, password, name string,
	age int8,
) InterfaceUserDomain {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

func NewUserUpdateDomain(
	name string,
	age int8,
) InterfaceUserDomain {
	return &userDomain{
		name: name,
		age:  age,
	}
}

func NewUserLoginDomain(
	email,
	password string,
) InterfaceUserDomain {
	return &userDomain{
		email:    email,
		password: password,
	}
}
