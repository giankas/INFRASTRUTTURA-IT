package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struttura Dominio
type Domain struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Simulazione di domini gestiti
var domains = []Domain{
	{ID: 1, Name: "esempio1.com"},
	{ID: 2, Name: "esempio2.net"},
}

func GetDomains(c *gin.Context) {
	c.JSON(http.StatusOK, domains)
}
