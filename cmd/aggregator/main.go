package main

import (
	"log"
	"sync"

	"github.com/729holdings/raajje-headlines/internal/database"
	"github.com/729holdings/raajje-headlines/internal/scraper"
)

func main() {
	database.InitDB()
	log.Println("🚀 729 Agentic Newsroom: Aggregator Cycle Started")

	// Run single cycle for Cloud Run Jobs
	runScraperCycle()
    
    log.Println("🏁 Aggregator Cycle Complete. Exiting gracefully.")
}

func runScraperCycle() {
	log.Println("🔄 Starting Concurrent Scrape & RSS Ingest Cycle...")
	
	// 1. Run RSS Ingestion
	scraper.IngestRSSSources()

	var wg sync.WaitGroup
	
	// List of supported sources
	sources := []func() ([]scraper.NewsArticle, error){
		scraper.ScrapeMihaaru,
		scraper.ScrapeSun,
		scraper.ScrapeVaguthu,
		scraper.ScrapeVNews,
	}

	for _, scrapeFunc := range sources {
		wg.Add(1)
		go func(f func() ([]scraper.NewsArticle, error)) {
			defer wg.Done()
			articles, err := f()
			if err != nil {
				log.Printf("Scrape error: %v", err)
				return
			}

			for _, a := range articles {
				// Check if already exists by URL
				var count int64
				database.DB.Model(&database.Article{}).Where("original_url = ?", a.Link).Count(&count)
				if count == 0 {
					article := database.Article{
						RawHeadline: a.Title,
						RawBody:     a.Body,
						OriginalURL: a.Link,
						SourceID:    1, // Placeholder for source mapping
						Status:      "pending_review",
					}
					if err := database.DB.Create(&article).Error; err != nil {
						log.Printf("❌ Failed to save article: %v", err)
					} else {
						log.Printf("📥 Saved new article: %s", a.Title)
					}
				} else {
					log.Printf("⏭️  Skipping existing article: %s", a.Title)
				}
			}
		}(scrapeFunc)
	}

	wg.Wait()
	log.Println("✅ Scrape cycle complete. All sources synced.")
}
