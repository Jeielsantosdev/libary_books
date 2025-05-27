package services

import (
	
	"time"

	"github.com/Jeielsantosdev/libary_books/config"
	
	"github.com/golang-jwt/jwt/v5"
)



func GenerateToken(username string)(string, error){
	claims := jwt.MapClaims{
		"username":username,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
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