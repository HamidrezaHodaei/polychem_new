package main

import (
	"database/sql"
	"log"
	"time"

	"polychem-auth/config"
	"polychem-auth/handlers"
	"polychem-auth/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// بارگذاری تنظیمات
	config.Load()

	// اتصال به دیتابیس
	db, err := sql.Open("mysql", config.AppConfig.GetDSN())
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// تست اتصال
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// تنظیمات Connection Pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	log.Println("✅ Connected to database successfully")

	// ایجاد Router
	if config.AppConfig.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// CORS Configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.AppConfig.FrontendURL, "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// ایجاد Auth Handler
	authHandler := handlers.NewAuthHandler(db)

	// Public Routes
	public := router.Group("/api")
	{
		public.POST("/login", authHandler.Login)
	}

	// Protected Routes (نیاز به توکن)
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(db))
	{
		protected.POST("/logout", authHandler.Logout)
		protected.GET("/me", func(c *gin.Context) {
			userID := c.GetInt("user_id")
			username := c.GetString("username")
			email := c.GetString("email")

			c.JSON(200, gin.H{
				"id":       userID,
				"username": username,
				"email":    email,
			})
		})
	}

	// Health Check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"version": "1.0.0",
		})
	})

	// شروع سرور
	port := config.AppConfig.ServerPort
	log.Printf("\n╔════════════════════════════════════════╗\n")
	log.Printf("║   🚀 Server Started Successfully       ║\n")
	log.Printf("╠════════════════════════════════════════╣\n")
	log.Printf("║ Port:        %s                     ║\n", port)
	log.Printf("║ Frontend:    %s    ║\n", config.AppConfig.FrontendURL)
	log.Printf("║ Environment: %-22s ║\n", config.AppConfig.Environment)
	log.Printf("║ Database:    %-22s ║\n", config.AppConfig.DBName)
	log.Printf("╚════════════════════════════════════════╝\n\n")

	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}