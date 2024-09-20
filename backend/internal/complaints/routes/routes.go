package route



import (
	"github.com/gin-gonic/gin"
	"github.com/wailbentafat/hakathon/backend/internal/complaints/provider"
	"database/sql"
	a"github.com/wailbentafat/hakathon/backend/core/middelware"
)

func Complainroute(r *gin.Engine ,db *sql.DB) {
	complaints.SetDB(db)
	r.GET("/get_complain",a.AdminMiddleware(db)  , complaints.Get_complent)
	r.POST("/add_complain",a.STUFFMiddleware(db)  ,complaints.Add_complain)
}