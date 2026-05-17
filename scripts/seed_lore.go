package main

import (
	"context"
	"log"
	"time"

	"github.com/729holdings/raajje-headlines/internal/database"
	"github.com/google/uuid"
)

func main() {
	database.InitDB()

	hotspots := []database.LoreHotspot{
		{
			ID:       uuid.New(),
			Name:     "The First Atoll Anchor",
			LoreData: "Oral tradition speaks of a crystalline anchor dropped from a sky-vessel during the Great Submergence. It is said to be the metaphysical core of the central atoll.",
			Category: "ORAL_HISTORY",
			IsGeminiVerified: true,
			LastAwakening: time.Now(),
		},
		{
			ID:       uuid.New(),
			Name:     "Endheri Vessel DNA",
			LoreData: "The black coral reefs near Guraidhoo contain traces of bioluminescent minerals identical to the 'Hull DNA' of the pre-republic fleet.",
			Category: "HULL_DNA",
			IsGeminiVerified: true,
			LastAwakening: time.Now(),
		},
		{
			ID:       uuid.New(),
			Name:     "Sovereign Lume Node #01",
			LoreData: "A high-frequency signal node discovered in the ruins of an ancient watchtower. It pulses in sync with the Raahi Abjad force of the current sentinel.",
			Category: "SYSTEM_LORE",
			IsGeminiVerified: true,
			LastAwakening: time.Now(),
		},
	}

	for _, h := range hotspots {
		if err := database.DB.Create(&h).Error; err != nil {
			log.Printf("Failed to seed hotspot %s: %v", h.Name, err)
		} else {
			log.Printf("Successfully seeded: %s", h.Name)
		}
	}
}
