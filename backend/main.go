package main

import (
	"github.com/giankas/INFRASTRUTTURA-IT/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Registrazione di tutte le rotte
	routes.RegisterRoutes(router)

	// Avvio del server sulla porta 8080
	router.Run(":8080")
}
