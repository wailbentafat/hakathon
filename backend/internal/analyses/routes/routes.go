package routess

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	a "github.com/wailbentafat/hakathon/backend/core/middelware"
	 "github.com/wailbentafat/hakathon/backend/internal/analyses/providers"
)

func Complainrou(r *gin.Engine ,db *sql.DB) {
	analyse.SetDB(db)
r.GET("/analyse",a.AdminMiddleware(db)  , analyse.CompAnalyse)
r.GET("/today",a.AdminMiddleware(db),analyse.Analyictoday)
r.GET("/datedata",a.AdminMiddleware(db),analyse.Dateanalytics)
}