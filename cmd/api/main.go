package main

import (
	"log"
	"os"

	"github.com/729holdings/raajje-headlines/internal/api"
	"github.com/729holdings/raajje-headlines/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize Database Connection
	database.InitDB()

	app := fiber.New(fiber.Config{
		AppName: "Raajjé HEADLINES v1.0",
	})

	// Enable CORS for dashboard local development
	app.Use(cors.New())

	// Admin Routes (Requires Auth) for Editorial Approval
	admin := app.Group("/admin", basicauth.New(basicauth.Config{
		Users: map[string]string{
			os.Getenv("ADMIN_USER"): os.Getenv("ADMIN_PASS"),
		},
	}))
	
	admin.Get("/pending", api.GetPendingArticles) // Fetch articles for the Dashboard
	admin.Post("/approve/:id", api.ApproveArticle)

	// Auth Routes
	apiV1 := app.Group("/api")
	apiV1.Post("/register", api.Register)
	app.Post("/api/login", api.Login)
	app.Post("/api/social-login", api.SocialLogin)

	// --- 📡 Branded WebApp & Dashboard Serving ---
	app.Static("/assets", "./web/assets") // Logo, Icons, etc
	app.Static("/", "./web/app.html")      // Public Livefeed (Root)
	app.Static("/dashboard", "./web/index.html") // Editorial Review

	// Public Routes for the Mobile App/Website
	apiV1.Get("/feed", api.GetFeed)
	apiV1.Get("/article/:id", api.GetArticleDetail)
	apiV1.Get("/search", api.SearchArticles)

	// Agentic AI Action Endpoints (Seyku Bro CAFA & Vault Integration)
	aiV1 := app.Group("/v1")
	aiV1.Post("/gemini/analyze", api.AnalyzeGemini)
	aiV1.Post("/nano-banana/generate", api.GenerateImagine)
	aiV1.Post("/veo/video", api.CueVeo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
