package proviauth

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wailbentafat/hakathon/backend/core/jwt"
)

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

func Register(c*gin.Context){
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existingEmail string
		err := db.QueryRow("SELECT email FROM staff WHERE email = ?", req.Email).Scan(&existingEmail)

		if err != nil && err != sql.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			log.Printf("Database error: %v", err)
			return
		}

		if existingEmail != "" {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}
        IsAdmin:=false
		
		query := `INSERT INTO staff (password, email, is_admin, first_name, last_name) 
		          VALUES (?, ?, ?, ?, ?)`
		_, err = db.Exec(query, req.Password, req.Email, IsAdmin, req.FirstName, req.LastName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			log.Printf("Failed to create user: %v", err)
			return
		}
        token,err:=jwt.GenerateJWT(req.Email)
		if err != nil {
			return
		}
		
		c.JSON(http.StatusCreated, gin.H{"token":token})
	}

