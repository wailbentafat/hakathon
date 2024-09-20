package stuff

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
var db *sql.DB
func SetDB(database *sql.DB) {
	db = database
}
func Add_stuff(c *gin.Context) {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		First_name string `json:"first_name"`
		Second_name string `json:"second_name"`
	}

	if err := c.BindJSON(&user); err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	email := user.Email
	password := user.Password
	First_name:= user.First_name
	Second_name:= user.Second_name

	if email == "" || password == "" ||Second_name==""||First_name==""{
		log.Println("Missing required fields")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to generate password hash: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	
	query := `INSERT INTO staff (email, password,First_name,last_name) VALUES (?, ?,?,?)`
	_, err = db.Exec(query, email, hash,First_name,Second_name)
	if err != nil {
		log.Printf("Failed to insert user into database: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}
	