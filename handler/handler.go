package handler

import "github.com/gin-gonic/gin"

func Listuser(ctx *gin.Context){
	ctx.JSON(200, gin.H{"msg":"to aprendendo",})
		
}