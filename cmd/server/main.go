package main

import (
    "log"
    "github.com/henriquelazzarino/gookshelf/config"
    "github.com/henriquelazzarino/gookshelf/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    // Carregar variáveis de ambiente
    config.LoadEnv()

    // Inicializar Firebase
    config.InitFirebase()

    r := gin.Default()
    routes.SetupRoutes(r)
    
    log.Printf("Starting server on port %s...", config.Port)
    r.Run(":" + config.Port)
}
