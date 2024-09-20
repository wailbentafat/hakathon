package complaints

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func Add_complain(c *gin.Context) {
	type complain struct {
		StuffID     int    `json:"stuff_id"`
		BankCard    string `json:"bank_card"`
		Category    string `json:"category"`
		Name        string `json:"name"`
		Location    string `json:"location"`
		PhoneNumber string `json:"phone_number"`
		Description string `json:"description"`
		Satisfied   bool   `json:"satisfied"`
	}
     c.get("email")
	var complainval complain
	if err := c.BindJSON(&complainval); err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields
	if complainval.StuffID == 0 || complainval.BankCard == "" ||
		complainval.Category == "" || complainval.Name == "" ||
		complainval.Location == "" || complainval.PhoneNumber == "" ||
		complainval.Description == "" {
		log.Println("Missing required fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	// Check if the category exists
	query := `SELECT name FROM categories WHERE name = ?`
	var categoryExists string
	err := db.QueryRow(query, complainval.Category).Scan(&categoryExists)

	if err == sql.ErrNoRows {
		// If category does not exist, insert it
		query = `INSERT INTO categories (name) VALUES (?)`
		if _, err = db.Exec(query, complainval.Category); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	query = `INSERT INTO complaints (bank_card, name, category, location, phone_number, description, staff_id, satisfied) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err = db.Exec(query, complainval.BankCard, complainval.Name, 
		complainval.Category, complainval.Location, 
		complainval.PhoneNumber, complainval.Description, 
		complainval.StuffID, complainval.Satisfied)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Complaint added successfully"})
}
