package middlewares

import (
	"net/http"
	"strings"

	"github.com/Jeielsantosdev/libary_books/services"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        authHeader := ctx.GetHeader("Authorization")
        if authHeader == "" {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "Token não enviado"})
            return
        }

        tokenParts := strings.Split(authHeader, " ")
        if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "Formato de token inválido"})
            return
        }

        claims, err := services.ValidateToken(tokenParts[1])
        if err != nil {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "Token inválido ou expirado"})
            return
        }

        useremail, ok := claims["useremail"].(string)
        if !ok {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "useremail não encontrado no token"})
            return
        }

        ctx.Set("user", useremail)
        ctx.Next()
    }
}