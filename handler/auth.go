package handler

import (
	"net/http"

	"github.com/Jeielsantosdev/libary_books/models"
	"github.com/Jeielsantosdev/libary_books/services"
	"github.com/Jeielsantosdev/libary_books/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func Login(ctx *gin.Context){
	var input models.Users

	if err := ctx.ShouldBindJSON(&input); err != nil{
		utils.RespondError(ctx, 404, "JSON invalido")
		return
	}

	// autenticacao de usuario
	var user models.Users
	if err := models.DB.Where("useremail = ?", input.Useremail).First(&user).Error; err != nil{
		utils.RespondError(ctx, http.StatusUnauthorized, "Usuario ou senha invalidos")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password ),[]byte(input.Password));err != nil{
		utils.RespondError(ctx, http.StatusUnauthorized, "Usuário ou senha inválidos")
		return
	}
	

	token, err := services.GenerateToken(user.Useremail)
	if err != nil{
		utils.RespondError(ctx, http.StatusInternalServerError, "erro ao gerar token")
		return
	}

	utils.RespondJSON(ctx, 200, gin.H{"Token":token})

}

func Protected(ctx *gin.Context){
	username, exists := ctx.Get("user")
	if !exists {
		utils.RespondError(ctx, http.StatusUnauthorized, "Usuário não encontrado no contexto")
        return
	}
	utils.RespondJSON(ctx, http.StatusOK, gin.H{"mensagem":"Bem-vindo!","user":username})


	
}