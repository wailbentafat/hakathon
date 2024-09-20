package proviauth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wailbentafat/hakathon/backend/core/jwt"
	"golang.org/x/crypto/bcrypt"
)
func CreateAdmin(c *gin.Context) {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&user); err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := user.Email
	password := user.Password

	if email == "" || password == "" {
		log.Println("Missing required fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to generate password hash: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	isadmin := true

	query := `INSERT INTO staff (email, password, is_admin,First_name,last_name) VALUES (?, ?, ?, ?, ?)`
	_, err = db.Exec(query, email, string(hash),  isadmin,"admin","admin")
	if err != nil {
		log.Printf("Failed to insert admin into database: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := jwt.GenerateJWT(email)
	if err != nil {
		log.Printf("Failed to generate JWT: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Admin registered successfully")
	c.JSON(http.StatusOK, gin.H{"token": token})
}
