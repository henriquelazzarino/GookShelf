package routes

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/henriquelazzarino/gookshelf/controllers"
)

// JwtMiddleware verifies the JWT token in the Authorization header
func JwtMiddleware(secret string) func(c *gin.Context) {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/login" {
			// Skip middleware for login route
			c.Next()
			return
		}

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " prefix if present
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil {
			if validationErr, ok := err.(*jwt.ValidationError); ok {
				// Handle specific validation errors
				if validationErr.Errors&jwt.ValidationErrorMalformed != 0 {
					// Provide detailed error messages for debugging (optional)
					fmt.Println(validationErr.Errors)
				}
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			} else {
				// Handle other errors (e.g., network issues)
				c.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}

		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Claims can be accessed from the token
		claims := token.Claims.(jwt.MapClaims)
		c.Set("userId", claims["sub"]) // Example: store user ID in context

		c.Next()
	}
}

func SetupRoutes(router *gin.Engine, secret string) {
	// Rota para autenticação
	router.POST("/login", controllers.Login)

	// Book routes - Aplicar middleware antes de definir as rotas
	bookRoutes := router.Group("/books", JwtMiddleware(secret))
	{
		bookRoutes.POST("", controllers.CreateBook)
		bookRoutes.GET("", controllers.GetAllBooks)
		bookRoutes.GET("/:id", controllers.GetBook)
		bookRoutes.PUT("/:id", controllers.UpdateBook)
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
	}

	// User routes - Aplicar middleware antes de definir as rotas
	userRoutes := router.Group("/users", JwtMiddleware(secret))
	{
		userRoutes.POST("", controllers.CreateUser)
		userRoutes.GET("", controllers.GetAllUsers)
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
		// Changed paths to avoid conflict
		userRoutes.GET("/book/:bookId/user/:userId", controllers.AddBookToUser)
		userRoutes.GET("/remove/book/:bookId/user/:userId", controllers.RemoveBookFromUser)
	}
}
