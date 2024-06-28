package routes

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/henriquelazzarino/gookshelf/controllers"
)

func JwtMiddleware(secret string, allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil {
			if validationErr, ok := err.(*jwt.ValidationError); ok {
				fmt.Println(validationErr.Errors)
			}
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		userRole := claims["role"].(string)
		userId := claims["sub"].(string)

		roleAllowed := false
		for _, role := range allowedRoles {
			if role == userRole {
				roleAllowed = true
				break
			}
		}

		if !roleAllowed {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Set("userId", userId)
		c.Set("userRole", userRole)

		c.Next()
	}
}

func SetupRoutes(router *gin.Engine, secret string) {
	router.POST("/login", controllers.Login)

	bookRoutes := router.Group("/books")
	{
		bookRoutes.POST("", JwtMiddleware(secret, "admin", "librarian"), controllers.CreateBook)
		bookRoutes.GET("", JwtMiddleware(secret, "admin", "librarian", "regular"), controllers.GetAllBooks)
		bookRoutes.GET("/:id", JwtMiddleware(secret, "admin", "librarian", "regular"), controllers.GetBook)
		bookRoutes.PUT("/:id", JwtMiddleware(secret, "admin", "librarian"), controllers.UpdateBook)
		bookRoutes.DELETE("/:id", JwtMiddleware(secret, "admin", "librarian"), controllers.DeleteBook)
	}

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", JwtMiddleware(secret, "admin"), controllers.CreateUser)
		userRoutes.GET("", JwtMiddleware(secret, "admin"), controllers.GetAllUsers)
		userRoutes.GET("/:id", JwtMiddleware(secret, "admin"), controllers.GetUser)
		userRoutes.PUT("/:id", JwtMiddleware(secret, "regular"), controllers.UpdateUser)
		userRoutes.DELETE("/:id", JwtMiddleware(secret, "admin", "regular"), controllers.DeleteUser)
	}
}
