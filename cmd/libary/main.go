package main

import (
	"github.com/Jeielsantosdev/libary_books/models"
	"github.com/Jeielsantosdev/libary_books/router"
	"github.com/gin-gonic/gin"
)





func main(){
	models.ConnectDB()
	routers := gin.Default()
	router.InitializerRouter(routers)

	routers.Run(":8080")
}