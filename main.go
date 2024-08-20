package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kauakirchner/first-go-project/src/config/database/mongodb"
	"github.com/kauakirchner/first-go-project/src/controller"
	"github.com/kauakirchner/first-go-project/src/controller/routes"
	"github.com/kauakirchner/first-go-project/src/model/repository"
	"github.com/kauakirchner/first-go-project/src/model/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := mongodb.NewMongoDbConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error=%s \n", err.Error())
		return
	}
	repo := repository.NewUserRepository(db)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
