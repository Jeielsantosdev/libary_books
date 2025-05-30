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

func ListBooks(ctx *gin.Context) {
    var books []models.Book

    // Retrieve userID from context
    userID, exists := ctx.Get("userID")
    if !exists {
        utils.RespondError(ctx, http.StatusUnauthorized, "Usuário não autenticado")
        return
    }

    // Assert userID to uint
    userIDVal, ok := userID.(uint)
    if !ok {
        utils.RespondError(ctx, http.StatusInternalServerError, "Erro interno: userID inválido")
        return
    }

    // Query books for the specific user
    if err := models.DB.Where("user_id = ?", userIDVal).Find(&books).Error; err != nil {
        log.Printf("Error listing books: %v", err)
        utils.RespondError(ctx, http.StatusInternalServerError, "Erro ao listar livros")
        return
    }

    // Return the list of books as JSON
    utils.RespondJSON(ctx, http.StatusOK, books)
}

func Verbook(ctx *gin.Context){
	var book models.Book
	

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

	id := ctx.Param("id")
	if err := models.DB.First(&book, "id = ? AND user_id = ?",id,userIDVal).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}
	ctx.JSON(200,book)

}

func UpdateBook(ctx *gin.Context) {
	var book models.Book

	// Retrieve userID from context
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

	id := ctx.Param("id")
	if err := models.DB.First(&book, "id = ? AND user_id = ?", id, userIDVal).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Livro não encontrado"})
		return
	}

	if err := ctx.ShouldBindJSON(&book); err != nil {
		utils.RespondError(ctx, http.StatusBadRequest, "JSON inválido")
		return
	}

	if err := models.DB.Save(&book).Error; err != nil {
		log.Printf("Error updating book: %v", err)
		utils.RespondError(ctx, http.StatusInternalServerError, "Erro ao atualizar livro")
		return
	}

	utils.RespondJSON(ctx, http.StatusOK, gin.H{"mensagem": "Livro atualizado com sucesso", "book": book})
}