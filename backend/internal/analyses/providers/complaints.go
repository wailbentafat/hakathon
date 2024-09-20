package analyse

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func CompAnalyse(c *gin.Context) {
	response := gin.H{}

	var totalComplaints int
	query := `SELECT COUNT(*) FROM complaints`
	if err := db.QueryRow(query).Scan(&totalComplaints); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total complaints"})
		return
	}
	response["total_complaints"] = totalComplaints

	categoryQuery := `SELECT category, COUNT(*) as complaint_count FROM complaints GROUP BY category`
	rows, err := db.Query(categoryQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get complaints by category"})
		return
	}
	defer rows.Close()

	categoryCounts := make(map[string]int)
	for rows.Next() {
		var category string
		var count int
		if err := rows.Scan(&category, &count); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan category rows"})
			return
		}
		categoryCounts[category] = count
	}
	response["complaints_by_category"] = categoryCounts

	var satisfiedCount, totalCount int
	satisfactionQuery := `SELECT SUM(CASE WHEN satisfied THEN 1 ELSE 0 END) as satisfied_count, COUNT(*) as total_count FROM complaints`
	if err := db.QueryRow(satisfactionQuery).Scan(&satisfiedCount, &totalCount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get satisfaction rate"})
		return
	}
	response["satisfaction_rate"] = map[string]int{
		"satisfied_count": satisfiedCount,
		"total_count":     totalCount,
	}

	staffQuery := `SELECT staff_id, COUNT(*) as complaint_count FROM complaints GROUP BY staff_id`
	staffRows, err := db.Query(staffQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get complaints by staff"})
		return
	}
	defer staffRows.Close()

	staffCounts := make(map[int]int)
	for staffRows.Next() {
		var staffID int
		var count int
		if err := staffRows.Scan(&staffID, &count); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan staff rows"})
			return
		}
		staffCounts[staffID] = count
	}
	response["complaints_by_staff"] = staffCounts

	var missingFieldsCount int
	missingFieldsQuery := `SELECT COUNT(*) FROM complaints WHERE bank_card IS NULL OR name IS NULL OR description IS NULL`
	if err := db.QueryRow(missingFieldsQuery).Scan(&missingFieldsCount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get count of missing fields"})
		return
	}
	response["missing_fields_count"] = missingFieldsCount

	c.JSON(http.StatusOK, response)
}
