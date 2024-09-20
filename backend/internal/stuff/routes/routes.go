package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wailbentafat/hakathon/backend/internal/stuff/provider"
	"database/sql"
	a"github.com/wailbentafat/hakathon/backend/core/middelware"
)

func StuffRoutes(r *gin.Engine ,db *sql.DB) {
	stuff.SetDB(db)
	// GetStuff swagger documentation.
//
// @Summary Get information about stuff.
// @Description Retrieve information about stuff.
// @Produce json
// @Success 200 {object} Response
// @Router /stuff [get]
	r.GET("/stuff",a.AdminMiddleware(db)  , stuff.Get_stuff)
	// AddStuff swagger documentation.
//
// @Summary Add new stuff.
// @Description Create a new stuff entry.
// @Produce json
// @Success 200 {object} Response
// @Router /addstuff [post]
	r.POST("/addstuff",a.AdminMiddleware(db)  , stuff.Add_stuff)
	r.DELETE("/deletestuff/:id",a.AdminMiddleware(db)  ,stuff.Delete_stuff)
}