package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wailbentafat/hakathon/backend/internal/stuff/provider"
	"database/sql"
	a"github.com/wailbentafat/hakathon/backend/core/middelware"
)

func StuffRoutes(r *gin.Engine ,db *sql.DB) {
	stuff.SetDB(db)
	r.GET("/stuff",a.AdminMiddleware(db)  , stuff.Get_stuff)
	r.POST("/addstuff",a.AdminMiddleware(db)  , stuff.Add_stuff)
	r.DELETE("/deletestuff/:id",a.AdminMiddleware(db)  ,stuff.Delete_stuff)
}