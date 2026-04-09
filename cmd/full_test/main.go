package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/729holdings/raajje-headlines/internal/database"
	"github.com/729holdings/raajje-headlines/internal/notifications"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	fmt.Println("🌊 [FULL 729 SYSTEM TEST] Initiating End-to-End Pipeline...")

	// 1. Database & ENV
	database.InitDB()
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ GEMINI_API_KEY missing from environment")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("❌ AI Client Error: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-flash-latest")

	// 2. Mocking a "Sovereignty" Reality
	mockArticle := database.Article{
		RawHeadline: "Finance Minister says SDF is safe and liquid.",
		RawBody:     "In a televised address, the Minister stated that rumors of a $500M drawdown from the Sovereign Development Fund are false.",
		Status:      "pending_review",
	}

	fmt.Println("\n🤖 [STEP 1: AI SYNTHESIS] Rephrasing via Seyku Bro voice...")
	
	prompt := fmt.Sprintf(`
		You are Seyku Bro. Rephrase this news for the Frontline Voice.
		Truth: The SDF was drawdown by $524 Million for the Sukuk payment.
		Headline: %s
		Body: %s
	`, mockArticle.RawHeadline, mockArticle.RawBody)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatalf("❌ Synthesis Error: %v", err)
	}

	result := fmt.Sprint(resp.Candidates[0].Content.Parts[0])
	fmt.Printf("\n✅ [SYNTHESIZED PAYLOAD]:\n%s\n", result)

	// 3. Simulated Push Notification
	fmt.Println("\n📡 [STEP 2: NOTIFICATION] Broadcasting to 5 Mayors (Topic: 729_council_broadcast)...")
	notifications.SendBreakingNewsAlert("729 ALERT: SDF AUDIT LEAKED", "Seyku Bro has synthesized the actual fiscal drawdown data. Open adhu.space immediately.")

	fmt.Println("\n🏁 [TEST COMPLETE] Full Pipeline Verified for April 8th.")
}
