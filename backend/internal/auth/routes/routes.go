package routauth

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	"github.com/wailbentafat/hakathon/backend/internal/auth/provider"

	"database/sql"
)
// AuthRoutes handles authentication routes.
// 
// This endpoint is used to log in.
// @Summary Logs in a user
// @Description Logs in a user with the provided credentials
// @Accept json
// @Produce json
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Success 200 {string} string "Successfully logged in"
// @Router /login [post]
func  AuthRoutes(db *sql.DB, r *gin.Engine) {
	proviauth.SetDB(db)
	r.POST("/login", proviauth.Login)

	// This endpoint is used to add an admin user.
	// @Summary Adds an admin user
	// @Description Creates a new admin user with the provided details
	// @Accept json
	// @Produce json
	// @Param username query string true "Username"
	// @Param password query string true "Password"
	// @Success 200 {string} string "Admin user created successfully"
	// @Router /addadmin [post]
	r.POST("/addadmin", proviauth.CreateAdmin)
	r.POST("/register",proviauth.Register)

	// This endpoint is used to serve the Swagger UI.
	// @Summary Serves the Swagger UI
	// @Description Serves the Swagger UI for API documentation
	// @Router /swagger/*any [get]
	r.GET("/swaagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}