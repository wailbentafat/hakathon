package stuff

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get_stuff(c *gin.Context) {
	query := `SELECT id, email, first_name, last_name FROM staff`
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("Failed to get staff: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var staffs []map[string]interface{}
	for rows.Next() {
		var id int
		var email string
		var firstName string
		var lastName string
		if err := rows.Scan(&id, &email, &firstName, &lastName); err != nil {
			log.Printf("Failed to scan row: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		staff := map[string]interface{}{
			"id":         id,
			"email":      email,
			"first_name": firstName,
			"last_name":  lastName,
		}
		staffs = append(staffs, staff)
	}
	c.JSON(http.StatusOK, staffs)
}
