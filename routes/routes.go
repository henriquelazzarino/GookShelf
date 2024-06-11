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

    // Rotas para editores
    publisherRoutes := router.Group("/publishers")
    {
        publisherRoutes.POST("", controllers.CreatePublisher)
        publisherRoutes.GET("", controllers.GetAllPublishers)
        publisherRoutes.GET("/:id", controllers.GetPublisher)
        publisherRoutes.PUT("/:id", controllers.UpdatePublisher)
        publisherRoutes.DELETE("/:id", controllers.DeletePublisher)
    }
}
