package complaints

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	log.Printf("Setting database connection for complaints provider")
	db = database
	log.Printf("Database connection set")
}

func Get_complent(c *gin.Context) {
	// Query to select all columns from the complaints table
	query := `SELECT * FROM complaints`
	log.Printf("Executing query: %s", query)
	rows, err := db.Query(query)
	if err != nil ||rows==nil{
		log.Printf("Failed to get complaints: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	defer rows.Close()
	query=`SELECT COUNT(*) FROM complaints;`
	var count int
	err=db.QueryRow(query).Scan(&count)
    print(count)
	log.Printf("Got %d rows", )
	var complaints []map[string]interface{}
	for rows.Next() {
		// Define variables to hold the scanned values
		var id int
		var bank_card, name, category, location, phone_number, description, email_address, bank_name, website_url, national_id_number, card_type, incident_date, merchant_name, merchant_registration string
		var staff_id int
		var satisfied bool
		var created_at string
		var transaction_amount float64
		var transaction_date string

		// Scan the row into variables
		log.Printf("Scanning row")
		if err := rows.Scan(&id, &bank_card, &name, &category, &location, &phone_number, &description, &staff_id, &satisfied, &created_at, 
			&email_address, &bank_name, &website_url, &national_id_number, &card_type, &incident_date, &transaction_amount, 
			&transaction_date, &merchant_name, &merchant_registration); err != nil {
			log.Printf("Failed to scan row: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Create a map for the complaint and append it to the slice
		complaint := map[string]interface{}{
			"id":                     id,
			"bank_card":              bank_card,
			"name":                   name,
			"category":               category,
			"location":               location,
			"phone_number":           phone_number,
			"description":            description,
			"staff_id":               staff_id,
			"satisfied":              satisfied,
			"created_at":             created_at,
			"email_address":          email_address,
			"bank_name":              bank_name,
			"website_url":            website_url,
			"national_id_number":     national_id_number,
			"card_type":              card_type,
			"incident_date":          incident_date,
			"transaction_amount":     transaction_amount,
			"transaction_date":       transaction_date,
			"merchant_name":          merchant_name,
			"merchant_registration":   merchant_registration,
		}
		complaints = append(complaints, complaint)
	}

	log.Printf("Sending response")
	c.JSON(http.StatusOK, complaints)
}
