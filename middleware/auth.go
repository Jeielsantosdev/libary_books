package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Jeielsantosdev/libary_books/config"
	//"github.com/Jeielsantosdev/libary_books/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        authHeader := ctx.GetHeader("Authorization")
        if authHeader == "" {
            ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"erro": "Token não enviado"})
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de assinatura inválido")
			}
			return config.SecretKey, nil
		})
       if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			ctx.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["user_id"] == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token malformado"})
			ctx.Abort()
			return
		}

		userID := uint(claims["user_id"].(float64))
		ctx.Set("userID", userID)
		ctx.Next()
	}
}
    
