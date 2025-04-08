package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struttura Ticket
type Ticket struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

// Archivio in memoria dei ticket
var tickets = []Ticket{
    {ID: 1, Title: "Ticket 1", Content: "Problema di esempio"},
}

func GetTickets(c *gin.Context) {
    c.JSON(http.StatusOK, tickets)
}

func CreateTicket(c *gin.Context) {
    var newTicket Ticket
    if err := c.BindJSON(&newTicket); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Richiesta non valida"})
        return
    }
    newTicket.ID = len(tickets) + 1
    tickets = append(tickets, newTicket)
    c.JSON(http.StatusOK, newTicket)
}
