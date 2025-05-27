package router

import (
	"github.com/Jeielsantosdev/libary_books/handler"
	"github.com/gin-gonic/gin"
)


func InitializerRouter(router *gin.Engine){
	users := router.Group("api/users")
	//book := router.Group("api/book")


	users.POST("/creteuser",handler.CreateUser)
	users.GET("/getuser/:id", handler.GetUser)
	users.GET("/listuser", handler.ListAllUsers)
	users.PUT("/updateuser/:id", handler.UpdateUser)
	users.DELETE("/deleteuser/:id", handler.DeleteUser)

}