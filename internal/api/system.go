package api

import (
	"github.com/729holdings/raajje-headlines/internal/database"
	"github.com/gofiber/fiber/v2"
)

type SovereigntyStats struct {
	CurrentCitizens int     `json:"current_citizens"`
	GoalCitizens    int     `json:"goal_citizens"`
	NodeCount       int     `json:"node_count"`
	LumePercentage  float64 `json:"lume_percentage"`
}

func GetSovereigntyStats(c *fiber.Context) error {
	var count int64
	database.DB.Model(&database.User{}).Count(&count)

	stats := SovereigntyStats{
		CurrentCitizens: int(count) + 4200, // Adding baseline for the "First Spark" simulation
		GoalCitizens:    10000,
		NodeCount:       84,
		LumePercentage:  98.2,
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   stats,
	})
}
