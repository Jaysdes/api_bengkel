package main

import (
	"api_bengkel/config"
	"api_bengkel/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	config.ConnectDB()

	// Inisialisasi router
	r := gin.Default()

	// ✅ Middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://127.0.0.1:8000", "http://localhost:8000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Register semua route
	routes.SetupRoutes(r)

	// Railway kasih PORT lewat environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" // fallback buat lokal development
	}

	log.Printf("🚀 Server running on port %s", port)
	r.Run(":" + port)
}
