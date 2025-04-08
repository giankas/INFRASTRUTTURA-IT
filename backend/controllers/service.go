package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Struttura Servizio Attivo
type ActiveService struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Details string `json:"details"`
}

// Simulazione dei servizi attivi per gli utenti
var activeServices = map[string][]ActiveService{
	"user": {
		{ID: 1, Name: "Hosting Base", Details: "Hosting attivo fino al 2025-12-31"},
		{ID: 2, Name: "Email Professionale", Details: "Email configurata e attiva"},
	},
}

func GetActiveServices(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.StandardClaims)
	services, ok := activeServices[claims.Subject]
	if !ok {
		services = []ActiveService{}
	}
	c.JSON(http.StatusOK, services)
}
