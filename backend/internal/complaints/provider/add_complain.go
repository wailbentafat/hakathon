package complaints

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type complain struct {
	Name                string  `json:"name"`
	ContactNumber       string  `json:"contact_number"`
	EmailAddress        string  `json:"email_address"`
	NationalIdNumber    string  `json:"national_id_number"`
	CardType            string  `json:"card_type"`
	CardLastFourDigits  string  `json:"card_last_four_digits"`
	BankName            string  `json:"bank_name"`
	MerchantName        string  `json:"merchant_name"`
	MerchantRegistration string  `json:"merchant_registration"`
	TransactionDate     string  `json:"transaction_date"`
	TransactionAmount   float64 `json:"transaction_amount"`
	IncidentDate        string  `json:"incident_date"`
	AtmLocation         string  `json:"atm_location"`
	WebsiteUrl          string  `json:"website_url"`
	CallType            string  `json:"call_type"`
	CallSummary         string  `json:"call_summary"`
}

func Add_complain(c *gin.Context) {
	var complainval complain
	if err := c.BindJSON(&complainval); err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"err": "no email set"})
		log.Printf("No email set")
		return
	}

	query := `SELECT id FROM staff WHERE email = ?`
	var staffID int
	err := db.QueryRow(query, email).Scan(&staffID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "problem getting staff ID"})
		log.Printf("Error retrieving staff ID: %v", err)
		return
	}

	// Validate required fields
	query = `SELECT name FROM categories WHERE name = ?`
	var categoryExists string
	err = db.QueryRow(query, complainval.CallType).Scan(&categoryExists)
	if err == sql.ErrNoRows {
		query = `INSERT INTO categories (name) VALUES (?)`
		if _, err = db.Exec(query, complainval.CallType); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	query = `INSERT INTO complaint (
		bank_card, 
		name, 
		category, 
		location, 
		phone_number, 
		description, 
		staff_id, 
		email_address, 
		bank_name, 
		website_url, 
		national_id_number, 
		card_type, 
		incident_date, 
		transaction_amount, 
		transaction_date, 
		merchant_name, 
		merchant_registration
	) VALUES (?, ?, ?, ?, ?, ?,  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := db.Exec(query,
		complainval.CardLastFourDigits, // bank_card
		complainval.Name,                // name
		complainval.CallType,            // category
		complainval.AtmLocation,         // location
		complainval.ContactNumber,       // phone_number
		complainval.CallSummary,         // description
		staffID,                         // staff_id (use the variable obtained from the query)
		complainval.EmailAddress,        // email_address
		complainval.BankName,            // bank_name
		complainval.WebsiteUrl,          // website_url
		complainval.NationalIdNumber,    // national_id_number
		complainval.CardType,            // card_type
		complainval.IncidentDate,        // incident_date
		complainval.TransactionAmount,    // transaction_amount
		complainval.TransactionDate,      // transaction_date
		complainval.MerchantName,       
		complainval.MerchantRegistration  ,
	)

	if err != nil {
		log.Printf("Failed to insert complaint: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println(res)
	c.JSON(http.StatusOK, gin.H{"message": "Complaint added successfully"})
}