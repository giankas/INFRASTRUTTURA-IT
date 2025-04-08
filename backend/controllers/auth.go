package controllers

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Chiave segreta per JWT (da proteggere in ambiente di produzione)
var jwtKey = []byte("my_secret_key")

// Struttura per le credenziali di login/registrazione
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Simulazione di un archivio utenti in memoria
var users = map[string]string{
	"user": "password", // Utente di esempio
}

func Login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Richiesta non valida"})
		return
	}
	// Verifica le credenziali
	if storedPwd, ok := users[creds.Username]; !ok || storedPwd != creds.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenziali non valide"})
		return
	}
	// Generazione del token JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   creds.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossibile generare il token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Register(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Richiesta non valida"})
		return
	}
	// Controlla se l'utente esiste già
	if _, exists := users[creds.Username]; exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Utente già esistente"})
		return
	}
	// Registra il nuovo utente
	users[creds.Username] = creds.Password
	c.JSON(http.StatusOK, gin.H{"message": "Registrazione avvenuta con successo"})
}
