package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type InterfaceUserDomain interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetJSONValue() (string, error)

	SetID(string)
	SetName(string)
	SetAge(int8)

	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) InterfaceUserDomain {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

type userDomain struct {
	Id       string
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}
func (ud *userDomain) GetPassword() string {
	return ud.Password
}
func (ud *userDomain) GetName() string {
	return ud.Name
}
func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func (ud *userDomain) SetID(id string) {
	ud.Id = id
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
