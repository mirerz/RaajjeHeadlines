package scraper

import (
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeVaguthu() ([]NewsArticle, error) {
	var articles []NewsArticle
	c := colly.NewCollector(
		colly.AllowedDomains("vaguthu.mv", "www.vaguthu.mv"),
	)

	// Selector for Vaguthu.mv - Targeting the news links
	c.OnHTML("a[href*='/news/'], a[href*='/viyafaari/']", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		
		articleCollector := c.Clone()
		articleCollector.OnHTML("article", func(ae *colly.HTMLElement) {
			// Vaguthu uses distinctive h2/h3 for headlines
			title := ae.ChildText("h1")
			if title == "" {
				title = ae.ChildText("h2")
			}

			article := NewsArticle{
				Source:    "Vaguthu",
				Title:     strings.TrimSpace(title),
				Link:      link,
				Timestamp: ae.ChildText(".summary li span"), // e.g. "3 ގަޑިއިރު ކުރިން"
				Body:      strings.TrimSpace(ae.ChildText(".single-content")),
			}
			
			if article.Title != "" && article.Body != "" {
				articles = append(articles, article)
			}
		})
		articleCollector.Visit(link)
	})

	err := c.Visit("https://vaguthu.mv/")
	if err != nil {
		log.Printf("Vaguthu visit error: %v", err)
	}
	return articles, err
}
