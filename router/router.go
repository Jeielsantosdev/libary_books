package router

import (
	"github.com/Jeielsantosdev/libary_books/handler"
	"github.com/Jeielsantosdev/libary_books/middleware"
	"github.com/gin-gonic/gin"
)

// 

func InitializerRouter(router *gin.Engine){
	libary := router.Group("api/")
	//book := router.Group("api/book")




	libary.POST("/creteuser",handler.CreateUser)
	libary.GET("/getuser/:id", handler.GetUser)
	libary.GET("/listuser", handler.ListAllUsers)
	

	libary.POST("/login", handler.Login)
	protected := router.Group("/api/user")
	protected.Use(middlewares.AuthMiddleware())
	{
		// Rotas de usuário
		protected.PUT("/updateuser/:id", handler.UpdateUser) 
		protected.DELETE("/deleteuser/:id", handler.DeleteUser)

		// Rotas protegidas
		protected.GET("/protected", handler.Protected)

		// Rotas de livros
		protected.POST("/book/create", handler.CreateBook)
		protected.GET("/book/list", handler.ListBooks)
		protected.GET("/book/:id", handler.Verbook)
		protected.PUT("/book/update/:id", handler.UpdateBook)
	}

	
	
}