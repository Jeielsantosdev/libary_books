package handler

import (
	"log"
	"net/http"

	"github.com/Jeielsantosdev/libary_books/models"
	"github.com/Jeielsantosdev/libary_books/utils"
	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "JSON inválido")
		return
	}

	// Validate required fields (example)
	if book.Title == "" {
		utils.RespondError(ctx, http.StatusBadRequest, "Título do livro é obrigatório")
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		utils.RespondError(ctx, http.StatusUnauthorized, "Usuário não autenticado")
		return
	}

	userIDVal, ok := userID.(uint)
	if !ok {
		utils.RespondError(ctx, http.StatusInternalServerError, "Erro interno: userID inválido")
		return
	}

	book.UserID = userIDVal
	if err := models.DB.Create(&book).Error; err != nil {
		// Log error (assuming a logger is available)
		log.Printf("Error creating book: %v", err)
		utils.RespondError(ctx, http.StatusInternalServerError, "Erro ao criar livro")
		return
	}

	utils.RespondJSON(ctx, http.StatusCreated, gin.H{"mensagem": "Livro criado com sucesso", "book": book})
}