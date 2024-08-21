package main

import (
	"github.com/kauakirchner/first-go-project/src/controller"
	"github.com/kauakirchner/first-go-project/src/model/repository"
	"github.com/kauakirchner/first-go-project/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(db *mongo.Database) controller.InterfaceUserController {
	repo := repository.NewUserRepository(db)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
