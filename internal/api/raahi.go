package api

import (
	"github.com/gofiber/fiber/v2"
)

// Abjad Map for Dhivehi/Arabic characters
var abjadMap = map[rune]int{
	'ا': 1, 'ބ': 2, 'ޖ': 3, 'ދ': 4, 'ހ': 5, 'ވ': 6, 'ޒ': 7, 'ޙ': 8, 'ޠ': 9,
	'ޔ': 10, 'ކ': 20, 'ލ': 30, 'މ': 40, 'ނ': 50, 'ސ': 60, 'ޢ': 70, 'ފ': 80, 'ޞ': 90,
	'ޤ': 100, 'ރ': 200, 'ޝ': 300, 'ތ': 400, 'ޘ': 500, 'ޚ': 600, 'ޛ': 700, 'ޟ': 800, 'ޡ': 900,
	'ޣ': 1000,
}

type RaahiRequest struct {
	Name string `json:"name"`
}

func CalculateAbjad(c *fiber.Ctx) error {
	var req RaahiRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	totalValue := 0
	for _, char := range req.Name {
		if val, ok := abjadMap[char]; ok {
			totalValue += val
		}
	}

	// The "Power of 7" Reduction Logic
	finalForce := totalValue
	for finalForce > 21 {
		finalForce = (finalForce % 10) + (finalForce / 10)
	}

	return c.JSON(fiber.Map{
		"total_value": totalValue,
		"final_force": finalForce,
		"status":      "success",
		"tier":        "free", // Logic for subscriber detection can be added later
	})
}
