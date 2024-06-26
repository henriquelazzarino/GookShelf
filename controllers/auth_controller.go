package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henriquelazzarino/gookshelf/models"
	"github.com/henriquelazzarino/gookshelf/services"
)

func Login(c *gin.Context) {
	var login models.Auth
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := services.Login(login.Email, login.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
