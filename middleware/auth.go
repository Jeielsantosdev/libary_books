package middleware

import (
	"net/http"
	"strings"

	"github.com/Jeielsantosdev/libary_books/services"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware()gin.HandlerFunc{
	return func (ctx *gin.Context)  {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == ""{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro":"token nao enviado"})
			return
		}

		tokenParts := strings.Split(authHeader, "")
		if len(tokenParts)!=2 || tokenParts[0]!= "Bearer"{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro":"Formato invalido"})
			return
		}

		claims, err := services.ValidateToken(tokenParts[1])
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro":"token invalido ou expirado"})
			return
		}

		ctx.Set("user",claims["username"])
		ctx.Next()
	}
}