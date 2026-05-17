package main

import (
	"context"
	"log"
	"os"

	"github.com/729holdings/raajje-headlines/internal/api"
	"github.com/729holdings/raajje-headlines/internal/ai"
	"github.com/729holdings/raajje-headlines/internal/database"
	"github.com/729holdings/raajje-headlines/internal/rewriter"
	"github.com/729holdings/raajje-headlines/internal/scraper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Initialize Database Connection
	database.InitDB()
	database.InitRedis()
	
	// Initialize Shared Gemini Client
	if err := ai.InitGemini(context.Background()); err != nil {
		log.Printf("Warning: Gemini initialization failed: %v\n", err)
	}

	// Start Agentic Rewriter in the background
	go rewriter.RunAgent(context.Background())

	// Start X-Intelligence Node (Monitoring handles)
	go scraper.ListenToXIntelligence()

	app := fiber.New(fiber.Config{
		AppName: "Raajjé HEADLINES v1.0",
	})

	// Enable CORS for dashboard local development
	app.Use(cors.New())

	// Initialize Firebase Auth
	if err := api.InitFirebase(); err != nil {
		log.Printf("Warning: Failed to initialize Firebase Auth: %v\n", err)
	}

	// Admin Routes (Requires Auth) for Editorial Approval
	// Replaced basicauth with Firebase Auth Middleware
	admin := app.Group("/admin", api.FirebaseAuthMiddleware())
	
	admin.Get("/pending", api.GetPendingArticles) // Fetch articles for the Dashboard
	admin.Post("/approve/:id", api.ApproveArticle)

	// Auth Routes
	apiV1 := app.Group("/api")
	apiV1.Post("/register", api.Register)
	app.Post("/api/login", api.Login)
	app.Post("/api/social-login", api.SocialLogin)

	// Public Routes for the Mobile App/Website
	apiV1.Get("/feed", api.GetFeed)
	apiV1.Get("/lore", api.GetLoreHotspots)
	apiV1.Get("/lore/search", api.SearchLore)
	apiV1.Get("/lore/:id", api.GetLoreDetail)
	apiV1.Get("/article/:id", api.GetArticleDetail)
	apiV1.Get("/search", api.SearchArticles)
	apiV1.Get("/sovereignty", api.GetSovereigntyStats)

	// Agentic AI Action Endpoints (Seyku Bro CAFA & Vault Integration)
	aiV1 := app.Group("/v1")
	aiV1.Post("/gemini/analyze", api.AnalyzeGemini)
	aiV1.Post("/nano-banana/generate", api.GenerateImagine)
	aiV1.Post("/veo/video", api.CueVeo)
	aiV1.Post("/raahi/calculate", api.CalculateAbjad)
	aiV1.Post("/synthesize", api.SynthesizeVoice)

	// --- 📡 Branded WebApp & Dashboard Serving ---
	app.Static("/assets", "./web/assets") 
	app.Static("/dashboard", "./web/index.html") 
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./web/app.html")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	log.Fatal(app.Listen(":" + port))
}
