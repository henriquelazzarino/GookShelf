package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henriquelazzarino/gookshelf/models"
	"github.com/henriquelazzarino/gookshelf/services"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := services.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func GetAllUsers(c *gin.Context) {
	authUserRole := c.GetString("userRole")
	if authUserRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have access to this resource"})
		return
	}

	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	authUserId := c.GetString("userId")
	authUserRole := c.GetString("userRole")

	if authUserId != id && authUserRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have access to this resource"})
		return
	}

	user, err := services.GetUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func AddBookToUser(c *gin.Context) {
	bookId := c.Param("bookId")
	userId := c.Param("userId")
	authUserId := c.GetString("userId")

	if authUserId != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have access to modify this user's books"})
		return
	}

	if err := services.AddBookToUser(userId, bookId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func RemoveBookFromUser(c *gin.Context) {
	bookId := c.Param("bookId")
	userId := c.Param("userId")
	authUserId := c.GetString("userId")

	if authUserId != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have access to modify this user's books"})
		return
	}

	if err := services.RemoveBookFromUser(userId, bookId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	authUserId := c.GetString("userId")

	if authUserId != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only update your own user"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateUser(id, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	authUserId := c.GetString("userId")
	authUserRole := c.GetString("userRole")

	if authUserRole != "admin" && authUserId != id {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can only delete your own user"})
		return
	}

	if err := services.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
