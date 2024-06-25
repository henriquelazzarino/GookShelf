package routes

import (
    "github.com/henriquelazzarino/gookshelf/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
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
