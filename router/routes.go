package router

import (
	"github.com/Jeielsantosdev/libary_books/handler"
	"github.com/gin-gonic/gin"
)

func Initrouter(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		v1.GET("/user", handler.Listuser )
	}

}