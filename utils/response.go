package utils

import "github.com/gin-gonic/gin"



func RespondJSON(ctx *gin.Context, status int, payload interface{}) {
	ctx.JSON(status, payload)
}

func RespondError(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, gin.H{"erro": message})
}