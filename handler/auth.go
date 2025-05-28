package handler

import (
    "net/http"

    "github.com/Jeielsantosdev/libary_books/models"
    "github.com/Jeielsantosdev/libary_books/services"
    "github.com/Jeielsantosdev/libary_books/utils"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
    var input models.Users
    if err := ctx.ShouldBindJSON(&input); err != nil {
        utils.RespondError(ctx, http.StatusBadRequest, "JSON inválido")
        return
    }

    // Validate input
    if input.Useremail == "" || input.Password == "" {
        utils.RespondError(ctx, http.StatusBadRequest, "Email ou senha não fornecidos")
        return
    }

    // Authenticate user
    var user models.Users
    if err := models.DB.Where("useremail = ?", input.Useremail).First(&user).Error; err != nil {
        utils.RespondError(ctx, http.StatusUnauthorized, "Credenciais inválidas")
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        utils.RespondError(ctx, http.StatusUnauthorized, "Credenciais inválidas")
        return
    }

    token, err := services.GenerateToken(user.ID)
    if err != nil {
        utils.RespondError(ctx, http.StatusInternalServerError, "Erro ao gerar token")
        return
    }

    utils.RespondJSON(ctx, http.StatusOK, gin.H{"token": token})
}

func Protected(ctx *gin.Context) {
    userID, exists := ctx.Get("userID")
    if !exists {
        utils.RespondError(ctx, http.StatusUnauthorized, "Usuário não autenticado")
        return
    }
    utils.RespondJSON(ctx, http.StatusOK, gin.H{"mensagem": "Bem-vindo!", "user_id": userID})
}