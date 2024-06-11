package controllers

import (
    "net/http"
    "github.com/henriquelazzarino/gookshelf/models"
	"github.com/henriquelazzarino/gookshelf/services"
    "github.com/gin-gonic/gin"
)

func CreatePublisher(c *gin.Context) {
    var publisher models.Publisher
    if err := c.ShouldBindJSON(&publisher); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := services.CreatePublisher(&publisher); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, publisher)
}

func GetAllPublishers(c *gin.Context) {
    publishers, err := services.GetAllPublishers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, publishers)
}

func GetPublisher(c *gin.Context) {
    id := c.Param("id")
    publisher, err := services.GetPublisher(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, publisher)
}

func UpdatePublisher(c *gin.Context) {
    id := c.Param("id")
    var publisher models.Publisher
    if err := c.ShouldBindJSON(&publisher); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := services.UpdatePublisher(id, &publisher); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, publisher)
}

func DeletePublisher(c *gin.Context) {
    id := c.Param("id")
    if err := services.DeletePublisher(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}
