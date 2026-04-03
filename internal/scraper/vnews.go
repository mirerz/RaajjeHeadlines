package scraper

import (
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeVNews() ([]NewsArticle, error) {
	var articles []NewsArticle
	c := colly.NewCollector(
		colly.AllowedDomains("vnews.mv", "www.vnews.mv"),
	)

	// Selector for Vnews.mv - Targeting the news card titles
	c.OnHTML("a.news-card", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		
		articleCollector := c.Clone()
		articleCollector.OnHTML(".article-page", func(ae *colly.HTMLElement) {
			title := ae.ChildText("h1")
			if title == "" {
				title = ae.ChildText(".article-title")
			}

			article := NewsArticle{
				Source:    "Vnews",
				Title:     strings.TrimSpace(title),
				Link:      link,
				Timestamp: ae.ChildText("strong"), // VNews often uses strong for date
				Body:      strings.TrimSpace(ae.ChildText(".article-body")),
			}
			
			if article.Title != "" && article.Body != "" {
				articles = append(articles, article)
			}
		})
		articleCollector.Visit(link)
	})

	err := c.Visit("https://vnews.mv/")
	if err != nil {
		log.Printf("Vnews visit error: %v", err)
	}
	return articles, err
}
