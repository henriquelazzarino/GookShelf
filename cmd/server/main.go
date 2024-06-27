package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/henriquelazzarino/gookshelf/config"
	"github.com/henriquelazzarino/gookshelf/models"
	"github.com/henriquelazzarino/gookshelf/routes"
	"github.com/henriquelazzarino/gookshelf/services"
)

func main() {
	// Carregar variáveis de ambiente
	config.LoadEnv()

	// Inicializar Firebase
	config.InitFirebase()

	r := gin.Default()
	routes.SetupRoutes(r, config.JWTSecret)

	log.Printf("Starting server on port %s...", config.Port)
	r.Run(":" + config.Port)

	InitUser()
}

func InitUser() {
	// Create a new user
	user := models.User{
		Name:     "Admin",
		Password: "password",
		Email:    "admin@admin.com",
		Age:      30,
	}

	services.CreateUser(&user)
}
