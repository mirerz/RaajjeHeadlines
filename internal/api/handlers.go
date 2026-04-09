package api

import (
	"encoding/json"
	"time"

	"github.com/729holdings/raajje-headlines/internal/database"
	"github.com/729holdings/raajje-headlines/internal/notifications"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetFeed returns the latest headlines (Cached via Redis)
func GetFeed(c *fiber.Ctx) error {
	const cacheKey = "news_feed_latest"

	// 1. Check Redis Cache
	cachedFeed, err := database.GetCache(cacheKey)
	if err == nil {
		var articles []database.Article
		if json.Unmarshal([]byte(cachedFeed), &articles) == nil {
			return c.JSON(articles)
		}
	}

	// 2. Cache Miss: Query Database
	var articles []database.Article
	result := database.DB.Select("id", "rephrased_headline_dv", "category", "is_breaking", "section", "news_type", "location", "created_at", "view_count").
		Where("status = ?", "published").
		Order("view_count desc, created_at desc").
		Limit(20).
		Find(&articles)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch news feed"})
	}

	// 3. Save to Redis (Expire in 5 Minutes for freshness)
	feedJSON, _ := json.Marshal(articles)
	database.SetCache(cacheKey, string(feedJSON), 5*time.Minute)

	return c.JSON(articles)
}

// GetPendingArticles returns articles that need editorial review
func GetPendingArticles(c *fiber.Ctx) error {
	var articles []database.Article
	result := database.DB.Where("status = ?", "pending_review").
		Order("created_at desc").
		Find(&articles)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not fetch pending articles"})
	}
	return c.JSON(articles)
}

// GetArticleDetail returns the full rephrased body (Triggered by "Read More")
func GetArticleDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	var article database.Article

	result := database.DB.First(&article, "id = ?", id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Article not found"})
	}

	// Increment view count for 729 Holdings analytics
	if err := database.DB.First(&article, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Article not found"})
	}

	// Increment View Count for priority ordering
	database.DB.Model(&article).UpdateColumn("view_count", gorm.Expr("view_count + ?", 1))

	// Access Control (Paywall Logic)
	isSubscribed := c.Get("X-User-Subscription") == "true"
	if article.IsPremium && !isSubscribed {
		// Return only partial content for non-subscribers
		return c.JSON(fiber.Map{
			"id":             article.ID,
			"headline":       article.RephrasedHeadlineDv,
			"summary":        article.SummaryEn,
			"body":           "This is premium content. Please subscribe to Raajjé HEADLINES to read the full story.",
			"is_premium":     true,
			"needs_upgrade":  true,
		})
	}

	return c.JSON(article)
}

func SearchArticles(c *fiber.Ctx) error {
	query := c.Query("q")
	var articles []database.Article
	database.DB.Where("status = ? AND (rephrased_headline_dv ILIKE ? OR rephrased_body_dv ILIKE ?)", 
		"published", "%"+query+"%", "%"+query+"%").
		Order("created_at desc").Limit(20).Find(&articles)
	return c.JSON(articles)
}

// ApproveArticle handles the editorial approval logic
func ApproveArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article database.Article

	if err := database.DB.First(&article, "id = ?", id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Article not found"})
	}

	// Capture the final edited content from the dashboard
	type ApprovalRequest struct {
		Headline            string `json:"headline"`
		Body                string `json:"body"`
		SubscriberBriefing  string `json:"subscriber_briefing"`
		SubscriberExplainer string `json:"subscriber_explainer"`
		VisualSource        string `json:"visual_source"`
		VisualURL           string `json:"visual_url"`
		IsPremium           bool   `json:"is_premium"`
		IsAd                bool   `json:"is_ad"`
		AdURL               string `json:"ad_url"`
	}
	var req ApprovalRequest
	if err := c.BodyParser(&req); err == nil && req.Headline != "" {
		article.RephrasedHeadlineDv = req.Headline
		article.RephrasedBodyDv = req.Body
		article.SubscriberBriefing = req.SubscriberBriefing
		article.SubscriberExplainer = req.SubscriberExplainer
		article.IsPremium = req.IsPremium
		article.IsAd = req.IsAd
		article.AdURL = req.AdURL
		
		// Check if editor actually changed anything
		if req.Headline != article.RephrasedHeadlineDv || req.Body != article.RephrasedBodyDv {
			article.IsCorrected = true
			article.OriginalAIRephrasing = article.RephrasedHeadlineDv + "\n\n" + article.RephrasedBodyDv
			article.RephrasedHeadlineDv = req.Headline
			article.RephrasedBodyDv = req.Body
		}
	}

	article.Status = "published"
	database.DB.Save(&article)

	// Invalidate Cache for fresh feed delivery
	database.ClearCache("news_feed_latest")

	if article.IsBreaking {
		// Fire-and-forget goroutine to keep the API responsive
		go notifications.SendBreakingNewsAlert(article.RephrasedHeadlineDv, article.SummaryEn)
	}

	return c.SendStatus(200)
}

// --- Agentic AI Endpoint Proxies ---

// AnalyzeGemini proxies requests to the ADK Gemini model
func AnalyzeGemini(c *fiber.Ctx) error {
	type Request struct {
		Context string `json:"context"`
		Pillar  string `json:"pillar"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid payload"})
	}

	// This is where real ADK Go 1.0 calling logic would reside.
	// We return a mock response format matching the UI expectations.
	return c.JSON(fiber.Map{
		"strategy": "Gemini System Hook: Detected Pillar " + req.Pillar + ". Inject a dramatic tension in the intro based on your current markdown.",
	})
}

// GenerateImagine handles Nano-Banana image generation integration
func GenerateImagine(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"image_url": "https://assets.raajjeheadlines.news/auto-generated-preview.jpg",
	})
}

// CueVeo handles Veo script-to-video integrations
func CueVeo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"video_url": "https://assets.raajjeheadlines.news/auto-generated-video.mp4",
		"status":    "processing",
	})
}
