package route



import (
	"github.com/gin-gonic/gin"
	"github.com/wailbentafat/hakathon/backend/internal/complaints/provider"
	"database/sql"
	a"github.com/wailbentafat/hakathon/backend/core/middelware"
)

func Complainroute(r *gin.Engine ,db *sql.DB) {
	complaints.SetDB(db)
	// GetComplain swagger documentation.
//
// @Summary Get a complaint.
// @Description Retrieve a complaint.
// @Produce json
// @Success 200 {object} Response
// @Router /get_complain [get]
	r.GET("/get_complain",a.AdminMiddleware(db)  , complaints.Get_complent)
	// AddComplain swagger documentation.
//
// @Summary Add a new complaint.
// @Description Create a new complaint.
// @Produce json
// @Success 200 {object} Response
// @Router /add_complain [post]
	r.POST("/add_complain",a.STUFFMiddleware(db)  ,complaints.Add_complain)
}