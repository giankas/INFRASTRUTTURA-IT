package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struttura Contatto
type ContactMessage struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func ContactForm(c *gin.Context) {
	var msg ContactMessage
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Richiesta non valida"})
		return
	}
	// In un'app reale si integrerebbe l'invio di email (es. SendGrid, Mailgun)
	// Qui stampiamo in console per simulazione
	c.JSON(http.StatusOK, gin.H{"message": "Messaggio inviato con successo", "data": msg})
}
