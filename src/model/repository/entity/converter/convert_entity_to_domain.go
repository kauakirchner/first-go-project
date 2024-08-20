package converter

import (
	"github.com/kauakirchner/first-go-project/src/model"
	"github.com/kauakirchner/first-go-project/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.InterfaceUserDomain {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)
	domain.SetID(entity.ID.Hex())
	return domain
}
