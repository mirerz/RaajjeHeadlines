package main

import (
	"log"
	"github.com/729holdings/raajje-headlines/internal/database"
)

func main() {
	database.InitDB()
	result := database.DB.Model(&database.Article{}).
		Where("status = ?", "pending_review").
		Limit(20).
		Update("status", "published")
	
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	log.Printf("Successfully published %d articles for the portal.", result.RowsAffected)
}
