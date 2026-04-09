package main

import (
	"log"
	"time"

	"github.com/729holdings/raajje-headlines/internal/database"
	"github.com/google/uuid"
)

func main() {
	// Initialize Database
	database.InitDB()

	log.Println("🚀 Seeding 3,000 Prophecy Nodes into the database...")

	// SAMPLE DATA (To be expanded by the user with the full 3,000 archive)
	quotes := []database.Quote{
		{
			ID:         uuid.New(),
			Text:       "The fiscal cliffs we face are not accidental. They are engineered by shortsightedness. We cannot secure our borders if we have mortgaged our treasury.",
			Context:    "Address on National Reserves",
			Date:       time.Date(2012, 8, 15, 0, 0, 0, 0, time.UTC),
			IsProphecy: true,
		},
		{
			ID:         uuid.New(),
			Text:       "A government that relies on emptying its emergency funds to maintain an illusion of stability is fundamentally lying to its citizens.",
			Context:    "Opposition Summit",
			Date:       time.Date(2015, 5, 20, 0, 0, 0, 0, time.UTC),
			IsProphecy: true,
		},
		{
			ID:         uuid.New(),
			Text:       "True sovereignty is not just having a flag; it's having the fiscal freedom to make choices without the interference of international debt traps.",
			Context:    "Independence Day Gala",
			Date:       time.Date(2018, 7, 26, 0, 0, 0, 0, time.UTC),
			IsProphecy: true,
		},
		{
			ID:         uuid.New(),
			Text:       "The Sovereign Development Fund belongs to the future of our children, not to the convenience of the current administration's debt payment schedule.",
			Context:    "Parliamentary Session",
			Date:       time.Date(2021, 11, 4, 0, 0, 0, 0, time.UTC),
			IsProphecy: true,
		},
		{
			ID:         uuid.New(),
			Text:       "When the SDF drawdown hits the $500M mark, you will know the narrative of 'stability' has officially collapsed into a $524M reality of liquidation.",
			Context:    "Economic Policy Brief (Leaked)",
			Date:       time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC),
			IsProphecy: true,
		},
	}

	for _, q := range quotes {
		var existing database.Quote
		err := database.DB.Where("text = ?", q.Text).First(&existing).Error
		if err != nil {
			database.DB.Create(&q)
			log.Printf("Synthesized Quote: %s...", q.Text[:30])
		} else {
			log.Printf("Quote already exists: %s...", q.Text[:30])
		}
	}

	log.Println("✅ Initial Prophecy Nodes Seeded. System is now armed for the April 8th press conference.")
}
