package controllers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Struttura Pacchetto di Servizi
type Package struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Simulazione dei pacchetti disponibili
var packages = []Package{
	{ID: 1, Name: "Hosting Base", Description: "Pacchetto hosting per piccoli siti", Price: 9.99},
	{ID: 2, Name: "Sito Web Personalizzato", Description: "Sviluppo sito web con design personalizzato", Price: 499.99},
}

// Restituisce i pacchetti disponibili
func GetPackages(c *gin.Context) {
	c.JSON(http.StatusOK, packages)
}

// Struttura Ordine
type Order struct {
	ID         int     `json:"id"`
	PackageID  int     `json:"package_id"`
	Username   string  `json:"username"`
	TotalPrice float64 `json:"total_price"`
}

// Archivio in memoria degli ordini
var orders = []Order{}

func CreateOrder(c *gin.Context) {
	var newOrder Order
	if err := c.BindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Richiesta non valida"})
		return
	}
	// In un'app reale si integrerebbe il pagamento tramite gateway (es. Stripe, PayPal)
	newOrder.ID = len(orders) + 1
	// Associa l'ordine all'utente autenticato (estratto dal token, ad es. dal subject)
	claims := c.MustGet("claims").(*jwt.StandardClaims)
	newOrder.Username = claims.Subject

	// Simula il calcolo del prezzo in base al pacchetto scelto
	for _, pkg := range packages {
		if pkg.ID == newOrder.PackageID {
			newOrder.TotalPrice = pkg.Price
			break
		}
	}
	orders = append(orders, newOrder)
	c.JSON(http.StatusOK, newOrder)
}

func GetOrders(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.StandardClaims)
	userOrders := []Order{}
	for _, order := range orders {
		if order.Username == claims.Subject {
			userOrders = append(userOrders, order)
		}
	}
	c.JSON(http.StatusOK, userOrders)
}
