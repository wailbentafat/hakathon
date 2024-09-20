package complaints

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get_complent(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	
	limit := 10 
	offset := 0 
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = l
		}
	}

	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = o
		}
	}

	query := `SELECT id, bank_card, name, category, location, phone_number, description, staff_id, satisfied, created_at FROM complaints LIMIT ? OFFSET ?`
	rows, err := db.Query(query, limit, offset)
	if err != nil {
		log.Printf("Failed to get complaints: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var complaints []map[string]interface{}
	for rows.Next() {
		var id int
		var bank_card string
		var name string
		var category string
		var location string
		var phone_number string
		var description string
		var staff_id int
		var satisfied bool
		var created_at string 

		if err := rows.Scan(&id, &bank_card, &name, &category, &location, &phone_number, &description, &staff_id, &satisfied, &created_at); err != nil {
			log.Printf("Failed to scan row: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		complaint := map[string]interface{}{
			"id":           id,
			"bank_card":    bank_card,
			"name":         name,
			"category":     category,
			"location":     location,
			"phone_number": phone_number,
			"description":  description,
			"staff_id":     staff_id,
			"satisfied":    satisfied,
			"created_at":   created_at,
		}
		complaints = append(complaints, complaint)
	}

	c.JSON(http.StatusOK, complaints)
}
