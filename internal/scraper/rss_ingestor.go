package scraper

import (
	"log"
	"time"
	"strings"

	"github.com/mmcdole/gofeed"
	"github.com/gocolly/colly/v2"
	"github.com/729holdings/raajje-headlines/internal/database"
	"github.com/729holdings/raajje-headlines/internal/utils"
)

type SourceConfig struct {
	Name     string
	URL      string
	FeedPath string
	Tier     int
	Lang     string
	IsGov    bool
}

func IngestRSSSources() {
	fp := gofeed.NewParser()
	sources := []SourceConfig{
		// Official: Maldives Government & SOEs
		{Name: "presidency.gov.mv", URL: "https://presidency.gov.mv", FeedPath: "/dv/news/rss", Tier: 0, Lang: "dv", IsGov: true},
		{Name: "sarukaaru.online", URL: "https://sarukaaru.online", FeedPath: "/feed", Tier: 0, Lang: "dv", IsGov: true}, // updated domain
		{Name: "psmnews.mv", URL: "https://psmnews.mv", FeedPath: "/feed", Tier: 1, Lang: "dv", IsGov: true},

		// Tier 1: High Priority (Maldives)
		{Name: "sun.mv", URL: "https://sun.mv", FeedPath: "/feed", Tier: 1, Lang: "dv"}, // changed from /rss to /feed
		{Name: "avas.mv", URL: "https://avas.mv", FeedPath: "/feed", Tier: 1, Lang: "dv"},
		{Name: "adhadhu.com", URL: "https://adhadhu.com", FeedPath: "/feed", Tier: 1, Lang: "dv"},
		{Name: "raajje.mv", URL: "https://raajje.mv", FeedPath: "/feed", Tier: 1, Lang: "dv"},
		{Name: "thepress.mv", URL: "https://thepress.mv", FeedPath: "/feed", Tier: 1, Lang: "dv"},
		{Name: "oneonline.mv", URL: "https://oneonline.mv", FeedPath: "/feed", Tier: 1, Lang: "dv"},
		{Name: "dhen.mv", URL: "https://dhen.mv", FeedPath: "/feed", Tier: 1, Lang: "dv"},

		// Tier 2: General
		{Name: "voice.mv", URL: "https://voice.mv", FeedPath: "/feed", Tier: 2, Lang: "dv"},
		{Name: "dhuvas.mv", URL: "https://dhuvas.mv", FeedPath: "/feed", Tier: 2, Lang: "dv"},
		{Name: "ras.mv", URL: "https://ras.mv", FeedPath: "/feed", Tier: 2, Lang: "dv"},
		{Name: "miadhu.mv", URL: "https://miadhu.mv", FeedPath: "/feed", Tier: 2, Lang: "dv"},

		// Global Intelligence Nodes (Tier 4: World)
		{Name: "Al Jazeera", URL: "https://www.aljazeera.com", FeedPath: "https://www.aljazeera.com/xml/rss/all.xml", Tier: 4, Lang: "en"},
		{Name: "BBC News", URL: "https://feeds.bbci.co.uk", FeedPath: "http://feeds.bbci.co.uk/news/world/rss.xml", Tier: 4, Lang: "en"},
		{Name: "CNN", URL: "http://rss.cnn.com", FeedPath: "http://rss.cnn.com/rss/edition_world.rss", Tier: 4, Lang: "en"},
		{Name: "Reuters", URL: "https://www.reutersagency.com", FeedPath: "https://www.reutersagency.com/feed/", Tier: 4, Lang: "en"},
	}

	for _, src := range sources {
		feedURL := src.FeedPath
		if !strings.HasPrefix(src.FeedPath, "http") {
			feedURL = src.URL + src.FeedPath
		}
		
		feed, err := fp.ParseURL(feedURL)
		if err != nil {
			log.Printf("Failed to parse feed for %s: %v", src.Name, err)
			continue
		}

		for _, item := range feed.Items {
			var count int64
			database.DB.Model(&database.Article{}).Where("original_url = ?", item.Link).Count(&count)
			if count > 0 {
				continue
			}
// 2. Scrape full body
fullBody := ScrapeFullBody(item.Link)
if fullBody == "" {
	fullBody = item.Description
}

// 🔒 [SECURITY] Sanitize HTML to prevent injection and AI hallucination
cleanBody := utils.SanitizeHTML(fullBody)

			article := database.Article{
				SourceName:       src.Name,
				SourceTier:       src.Tier,
				OriginalURL:      item.Link,
				RawHeadline:      utils.NormalizeThaana(item.Title),
				RawBody:          utils.NormalizeThaana(cleanBody),
				IsVerifiedGov:    src.IsGov,
				OriginalLanguage: src.Lang,
				Status:           "pending_review",
			}
			
			if item.Image != nil {
				article.VisualURL = item.Image.URL
			}

			database.DB.Create(&article)
			log.Printf("📥 [%s] Ingested: %s", src.Name, item.Title)
		}
	}
}

func ScrapeFullBody(url string) string {
	var body string
	c := colly.NewCollector()
	
	// Diverse selectors for Maldivian and Global sites
	selectors := []string{
		"article", 
		".article-content", 
		".story-content", 
		".entry-content", 
		".article-body", 
		".post-content",
		".td-post-content",
		"main .content",
	}

	for _, selector := range selectors {
		c.OnHTML(selector, func(e *colly.HTMLElement) {
			if body == "" {
				body = strings.TrimSpace(e.Text)
			}
		})
	}

	c.SetRequestTimeout(15 * time.Second)
	c.Visit(url)
	return body
}
