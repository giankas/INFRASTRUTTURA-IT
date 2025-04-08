package routes

import (
	"project/controllers"
	"project/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Rotte di autenticazione
	router.POST("/api/login", controllers.Login)
	router.POST("/api/register", controllers.Register)

	// Rotte pubbliche per visualizzare informazioni
	router.GET("/api/packages", controllers.GetPackages)
	router.GET("/api/domains", controllers.GetDomains)

	// Rotte per il form di contatto
	router.POST("/api/contact", controllers.ContactForm)

	// Rotte protette (richiedono token JWT)
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Ticket di assistenza
		protected.GET("/tickets", controllers.GetTickets)
		protected.POST("/tickets", controllers.CreateTicket)

		// Ordini e servizi attivi
		protected.POST("/orders", controllers.CreateOrder)
		protected.GET("/orders", controllers.GetOrders)
		protected.GET("/services", controllers.GetActiveServices)
	}
}
