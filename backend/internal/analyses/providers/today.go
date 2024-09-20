package analyse

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Analyictoday(c *gin.Context) {
	response := gin.H{}

	today := time.Now().Format("2006-01-02") // Get today's date in the proper format

	// Total complaints today
	var totalComplaints int
	query := `SELECT COUNT(*) FROM complaints WHERE DATE(created_at) = ?`
	if err := db.QueryRow(query, today).Scan(&totalComplaints); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total complaints"})
		return
	}
	response["total_complaints"] = totalComplaints

	// Complaints by category today
	categoryQuery := `SELECT category, COUNT(*) as complaint_count FROM complaints WHERE DATE(created_at) = ? GROUP BY category`
	rows, err := db.Query(categoryQuery, today)
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

	// Satisfaction rate today
	var satisfiedCount, totalCount int
	satisfactionQuery := `SELECT SUM(CASE WHEN satisfied THEN 1 ELSE 0 END) as satisfied_count, COUNT(*) as total_count FROM complaints WHERE DATE(created_at) = ?`
	if err := db.QueryRow(satisfactionQuery, today).Scan(&satisfiedCount, &totalCount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get satisfaction rate"})
		return
	}
	response["satisfaction_rate"] = map[string]int{
		"satisfied_count": satisfiedCount,
		"total_count":     totalCount,
	}

	// Complaints by staff today
	staffQuery := `SELECT staff_id, COUNT(*) as complaint_count FROM complaints WHERE DATE(created_at) = ? GROUP BY staff_id`
	staffRows, err := db.Query(staffQuery, today)
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

	// Count of missing fields today
	var missingFieldsCount int
	missingFieldsQuery := `SELECT COUNT(*) FROM complaints WHERE (bank_card IS NULL OR name IS NULL OR description IS NULL) AND DATE(created_at) = ?`
	if err := db.QueryRow(missingFieldsQuery, today).Scan(&missingFieldsCount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get count of missing fields"})
		return
	}
	response["missing_fields_count"] = missingFieldsCount

	// Retrieve all complaints today
	var allComplaints []Complaint
	complaintsQuery := `SELECT * FROM complaints WHERE DATE(created_at) = ?`
	complaintsRows, err := db.Query(complaintsQuery, today)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get complaints"})
		return
	}
	defer complaintsRows.Close()

	for complaintsRows.Next() {
		var complaint Complaint
		if err := complaintsRows.Scan(&complaint.ID, &complaint.BankCard, &complaint.Name, &complaint.Category, &complaint.Location, &complaint.PhoneNumber, &complaint.Description, &complaint.StaffID, &complaint.Satisfied, &complaint.CreatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan complaint"})
			return
		}
		allComplaints = append(allComplaints, complaint)
	}
	response["complaints_today"] = allComplaints

	c.JSON(http.StatusOK, response)
}
