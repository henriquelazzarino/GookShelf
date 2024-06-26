package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/henriquelazzarino/gookshelf/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// Rota para autenticação
	router.POST("/login", controllers.Login)

	// Rotas para livros
	bookRoutes := router.Group("/books")
	{
		bookRoutes.POST("", controllers.CreateBook)
		bookRoutes.GET("", controllers.GetAllBooks)
		bookRoutes.GET("/:id", controllers.GetBook)
		bookRoutes.PUT("/:id", controllers.UpdateBook)
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
	}

	// Rotas para usuários
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", controllers.CreateUser)
		userRoutes.GET("", controllers.GetAllUsers)
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
	}
}
