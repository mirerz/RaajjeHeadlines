package scraper

import (
	"strings"

	"github.com/729holdings/raajje-headlines/internal/utils"
	"github.com/gocolly/colly/v2"
)

type NewsArticle struct {
	Source    string
	Title     string
	Link      string
	Timestamp string
	Body      string
}

func ScrapeMihaaru() ([]NewsArticle, error) {
	var articles []NewsArticle
	c := colly.NewCollector(
		colly.AllowedDomains("mihaaru.com", "www.mihaaru.com"),
	)

	// 1. Find all article links on the homepage/category page
	c.OnHTML("a.story-link", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		
		// Create a new collector to visit each article's full page
		articleCollector := c.Clone()
		articleCollector.OnHTML("article", func(ae *colly.HTMLElement) {
			article := NewsArticle{
				Source:    "Mihaaru",
				Title:     utils.NormalizeThaana(strings.TrimSpace(ae.ChildText("h1.story-title"))),
				Link:      link,
				Timestamp: ae.ChildAttr("time", "datetime"),
				Body:      utils.NormalizeThaana(strings.TrimSpace(ae.ChildText(".story-content"))),
			}
			articles = append(articles, article)
		})
		articleCollector.Visit(link)
	})

	err := c.Visit("https://mihaaru.com/local")
	return articles, err
}
