package services

import (
	
	"time"

	"github.com/Jeielsantosdev/libary_books/config"
	
	"github.com/golang-jwt/jwt/v5"
)



func GenerateToken(userID uint)(string, error){
	claims := jwt.MapClaims{
		"user_id":userID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SecretKey)
}

func ValidateToken(tokenString string)(jwt.MapClaims, error){
	token, err := jwt.Parse(tokenString, func (token *jwt.Token)(interface{},error){
		return config.SecretKey, nil
	} )
	if err != nil || !token.Valid{
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok{
		return nil, err
	}
	return claims, nil
}