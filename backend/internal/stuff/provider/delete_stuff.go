package stuff

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete_stuff(c *gin.Context) {
	id := c.Param("id")
    if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing ID"})
		return
	}

	query := `DELETE FROM staff WHERE id = ?`
	_, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Failed to delete stuff: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Stuff deleted successfully"})
}