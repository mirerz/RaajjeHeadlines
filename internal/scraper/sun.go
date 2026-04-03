package scraper

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeSun() ([]NewsArticle, error) {
	var articles []NewsArticle
	c := colly.NewCollector(
		colly.AllowedDomains("sun.mv", "www.sun.mv"),
	)

	// Selector for Sun.mv - Targeting the main news grid
	c.OnHTML(".news-title a", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		
		articleCollector := c.Clone()
		articleCollector.OnHTML(".article-content", func(ae *colly.HTMLElement) {
			article := NewsArticle{
				Source:    "Sun",
				Title:     strings.TrimSpace(ae.ChildText("h1.article-title")),
				Link:      link,
				Timestamp: ae.ChildAttr(".article-date time", "datetime"),
				Body:      strings.TrimSpace(ae.ChildText(".article-body")),
			}
			articles = append(articles, article)
		})
		articleCollector.Visit(link)
	})

	err := c.Visit("https://sun.mv/local")
	return articles, err
}
