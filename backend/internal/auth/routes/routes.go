package routauth

import (
	"github.com/gin-gonic/gin"
	"github.com/wailbentafat/hakathon/backend/internal/auth/provider"

	"database/sql"

)

func  AuthRoutes(db *sql.DB, r *gin.Engine) {
	proviauth.SetDB(db)
	r.POST("/login", proviauth.Login)
	r.POST("/addadmin", proviauth.CreateAdmin)
}