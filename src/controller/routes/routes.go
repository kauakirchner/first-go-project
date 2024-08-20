package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/first-go-project/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.InterfaceUserController) {
	r.GET("/user/:userId", userController.FindUserById)
	r.GET("/getUserByEmail/:email", userController.FindUserByEmail)
	r.POST("/user", userController.CreateUser)
	r.PUT("/user/:userId", userController.UpdateUser)
	r.DELETE("/user/:userId", userController.DeleteUser)
}
