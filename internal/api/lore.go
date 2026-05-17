package api

import (
	"github.com/729holdings/raajje-headlines/internal/ai"
	"github.com/729holdings/raajje-headlines/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetLoreHotspots retrieves all lore hotspots from the database
func GetLoreHotspots(c *fiber.Ctx) error {
	var hotspots []database.LoreHotspot
	result := database.DB.Find(&hotspots)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch lore data"})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   hotspots,
	})
}

// GetLoreDetail retrieves a specific lore hotspot by ID
func GetLoreDetail(c *fiber.Ctx) error {
	id := c.Params("id")
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid lore ID"})
	}

	var hotspot database.LoreHotspot
	result := database.DB.First(&hotspot, "id = ?", parsedID)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Lore hotspot not found"})
	}

	// Phase 3: RAG Refinement (Gemini Rephrasing for Subscribers)
	isSubscribed := c.Get("X-User-Subscription") == "true"
	if isSubscribed {
		rephrased, err := ai.RephraseLore(c.Context(), hotspot.Name, hotspot.LoreData)
		if err == nil {
			hotspot.LoreData = rephrased // Inject the AI synthesis
		}
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   hotspot,
	})
}

// SearchLore searches hotspots by keyword
func SearchLore(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Query parameter 'q' is required"})
	}

	var hotspots []database.LoreHotspot
	result := database.DB.Where("name LIKE ? OR lore_data LIKE ?", "%"+query+"%", "%"+query+"%").Find(&hotspots)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Search failed"})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   hotspots,
	})
}
