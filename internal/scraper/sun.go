package scraper

import (
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeSun() ([]NewsArticle, error) {
	var articles []NewsArticle
	c := colly.NewCollector(
		colly.AllowedDomains("sun.mv", "www.sun.mv"),
	)

	// Selector for Sun.mv - Targeting the main news grid
	c.OnHTML("a.home-v2-news-thumb, a.home-v2-feetha-item, a.home-v2-featured", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		
		log.Printf("[Sun] Found link: %s", link)
		
		articleCollector := c.Clone()
		articleCollector.OnHTML("body", func(ae *colly.HTMLElement) {
			article := NewsArticle{
				Source:    "Sun",
				Title:     strings.TrimSpace(ae.ChildText(".article-title")),
				Link:      link,
				Timestamp: ae.ChildAttr(".article-date time", "datetime"),
				Body:      strings.TrimSpace(ae.ChildText(".article-body")),
			}
			if article.Title != "" {
				log.Printf("[Sun] Scraped article: %s", article.Title)
				articles = append(articles, article)
			} else {
				// Try fallback selectors
				article.Title = strings.TrimSpace(ae.ChildText("h1"))
				if article.Title != "" {
					log.Printf("[Sun] Scraped article (fallback): %s", article.Title)
					articles = append(articles, article)
				}
			}
		})
		articleCollector.Visit(link)
	})

	log.Println("[Sun] Visiting: https://sun.mv")
	err := c.Visit("https://sun.mv")
	return articles, err
}
