package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/first-go-project/src/controller"
	"github.com/kauakirchner/first-go-project/src/model"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(r *gin.RouterGroup, userController controller.InterfaceUserController) {
	r.GET("/user/:userId", model.AuthMiddleware, userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", model.AuthMiddleware, userController.FindUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:userId", model.AuthMiddleware, userController.UpdateUser)
	r.DELETE("/user/:userId", model.AuthMiddleware, userController.DeleteUser)
	r.POST("/login", userController.LoginUser)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
