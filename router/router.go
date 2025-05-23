package router

import (
	
	"github.com/gin-gonic/gin"
)


func Init(){
	// initialazer router
	router := gin.Default()
	//initializer routes
	Initrouter(router)
	//run the server :8080
	router.Run()
}