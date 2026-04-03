package main

import (
	"log"
	"sync"
	"time"

	"github.com/729holdings/raajje-headlines/internal/database"
	"github.com/729holdings/raajje-headlines/internal/scraper"
)

func main() {
	database.InitDB()
	log.Println("🚀 729 Agentic Newsroom: Aggregator Loop Started (15m Cycles)")

	// Run initial cycle
	runScraperCycle()

	// Ticker for periodic updates
	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			runScraperCycle()
		}
	}
}

func runScraperCycle() {
	log.Println("🔄 Starting Concurrent Scrape Cycle...")
	var wg sync.WaitGroup
	
	// List of supported sources
	sources := []func(){
		scraper.ScrapeMihaaru,
		scraper.ScrapeSun,
		scraper.ScrapeVaguthu,
		scraper.ScrapeVnews,
	}

	for _, scrapeFunc := range sources {
		wg.Add(1)
		go func(f func()) {
			defer wg.Done()
			f()
		}(scrapeFunc)
	}

	wg.Wait()
	log.Println("✅ Scrape cycle complete. All sources synced.")
}
