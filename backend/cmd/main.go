package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wailbentafat/hakathon/backend/internal/auth/routes"
	"github.com/wailbentafat/hakathon/backend/internal/db"
	"github.com/wailbentafat/hakathon/backend/internal/stuff/routes"
	"github.com/wailbentafat/hakathon/backend/internal/complaints/routes"
	"github.com/wailbentafat/hakathon/backend/internal/analyses/routes"



)




func main() {
    log.Println("Starting server")
    corsConfig := cors.Config{
        AllowOrigins:     []string{"*"}, 
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: false,
        MaxAge:           12 * time.Hour,
    }
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
    dsn := os.Getenv("DSN")
    log.Printf("Using database connection string: %s", dsn)
    

    

    database, err := db.InitDb(dsn)
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer func() {
        log.Println("Closing database connection")
        if err := database.Close(); err != nil {
            log.Fatalf("Failed to close database connection: %v", err)
        }
    }()
    router := gin.Default()
    router.Use(cors.New(corsConfig))
    routauth.AuthRoutes(database, router)
	routes.StuffRoutes(router,database)
	route.Complainroute(router,database)
	routess.Complainrou(router,database)
   

    err = http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Server started")
    }

