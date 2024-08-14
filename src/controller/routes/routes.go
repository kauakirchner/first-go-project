package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kauakirchner/first-go-project/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/user/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:email", controller.FindUserByEmail)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/:userId", controller.UpdateUser)
	r.DELETE("/user/:userId", controller.DeleteUser)
}
