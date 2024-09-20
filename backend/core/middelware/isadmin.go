package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	"github.com/wailbentafat/hakathon/backend/core/jwt"

	"database/sql")

func AdminMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		
		tokenStr := parts[1]
		claims, err := jwt.Parsejwt(tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
        query:=`SELECT is_admin FROM staff  WHERE email = ?`
		var isadmin bool
		err = db.QueryRow(query, claims).Scan(&isadmin)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "failed to check admin status"})
			c.Abort()
			return
		}

		if !isadmin {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		
		c.Set("email", claims)
        c.Set("isadmin", isadmin)
		
		c.Next()
	}
}
