package analyse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Complaint struct {
	ID          uint64  `json:"id"`
	BankCard    string  `json:"bank_card"`
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Location    string  `json:"location"`
	PhoneNumber string  `json:"phone_number"`
	Description string  `json:"description"`
	StaffID     *uint64 `json:"staff_id"` // Pointer to allow NULL values
	Satisfied   bool    `json:"satisfied"`
	CreatedAt   string  `json:"created_at"` // Consider using time.Time for better date handling
}

func Dateanalytics(c *gin.Context) {
	response := gin.H{}

	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	// Total complaints in the specified date range
	var totalComplaints int
	query := `SELECT COUNT(*) FROM complaints WHERE created_at BETWEEN ? AND ?`
	if err := db.QueryRow(query, startDate, endDate).Scan(&totalComplaints); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total complaints"})
		return
	}
	response["total_complaints"] = totalComplaints

	// Complaints by category in the specified date range
	categoryQuery := `SELECT category, COUNT(*) as complaint_count FROM complaints WHERE created_at BETWEEN ? AND ? GROUP BY category`
	rows, err := db.Query(categoryQuery, startDate, endDate)
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

	// Satisfaction rate in the specified date range
	var satisfiedCount, totalCount int
	satisfactionQuery := `SELECT SUM(CASE WHEN satisfied THEN 1 ELSE 0 END) as satisfied_count, COUNT(*) as total_count FROM complaints WHERE created_at BETWEEN ? AND ?`
	if err := db.QueryRow(satisfactionQuery, startDate, endDate).Scan(&satisfiedCount, &totalCount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get satisfaction rate"})
		return
	}
	response["satisfaction_rate"] = map[string]int{
		"satisfied_count": satisfiedCount,
		"total_count":     totalCount,
	}

	// Complaints by staff in the specified date range
	staffQuery := `SELECT staff_id, COUNT(*) as complaint_count FROM complaints WHERE created_at BETWEEN ? AND ? GROUP BY staff_id`
	staffRows, err := db.Query(staffQuery, startDate, endDate)
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

	// Missing fields count in the specified date range
	var missingFieldsCount int
	missingFieldsQuery := `SELECT COUNT(*) FROM complaints WHERE (bank_card IS NULL OR name IS NULL OR description IS NULL) AND created_at BETWEEN ? AND ?`
	if err := db.QueryRow(missingFieldsQuery, startDate, endDate).Scan(&missingFieldsCount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get count of missing fields"})
		return
	}
	response["missing_fields_count"] = missingFieldsCount

	// Retrieve all complaints in the specified date range
	var allComplaints []Complaint
	complaintsQuery := `SELECT id, bank_card, name, category, location, phone_number, description, staff_id, satisfied, created_at FROM complaints WHERE created_at BETWEEN ? AND ?`
	complaintsRows, err := db.Query(complaintsQuery, startDate, endDate)
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
	response["complaints"] = allComplaints

	c.JSON(http.StatusOK, response)
}
