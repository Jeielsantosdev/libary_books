package handler

import (
	"net/http"
	
	"strconv"

	"github.com/Jeielsantosdev/libary_books/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)


func CreateUser(ctx *gin.Context){
	var user models.Users

	if err:= ctx.ShouldBindJSON(&user); err !=nil{
		ctx.JSON(404,gin.H{"erro":err.Error()})
		return
	}
	

	//hash da senha
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro":"Erro ao gerar hash da senha"})
		return
	}

	user.Password = string(hash)
	// cria o users no db
	if err := models.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro":"erro ao salva o usuario no banco"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"mensagem":"Usuario criado com sucesso",
										"Usuario":user,	
									})
}
func ListAllUsers(ctx *gin.Context) {
	var users []models.Users
	if err := models.DB.Find(&users).Error; err != nil {
	    ctx.JSON(500, gin.H{"error": "Erro ao buscar usuários"})
	    return
	}
	if len(users) == 0 {
	    ctx.JSON(404, gin.H{"error": "Nenhum usuário encontrado"})
	    return
	}
ctx.JSON(200, users)

}

func GetUser(ctx *gin.Context) {
	idStr := ctx.Param("id")

	// Verifica se veio algo
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID não informado"})
		return
	}

	// Converte string para int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	var user models.Users
	if err := models.DB.First(&user, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"erro": "Usuário não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func UpdateUser(ctx *gin.Context){
	var user models.Users
	id := ctx.Param("id")
	//TODO:arruma o hash de senha ao fazer o update
	if err := models.DB.First(&user, id).Error; err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"erro":err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"erro":err.Error()})
		return
	}

	models.DB.Save(&user)
	ctx.JSON(200, user)
}

func DeleteUser(ctx *gin.Context){
	
	id := ctx.Param("id")

	if err := models.DB.Delete(&models.Users{},id).Error; err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro":"Erro ao deletar"})
		return
	}
	ctx.JSON(200, gin.H{"message":"Usuario deletado com sucesso"})
}
