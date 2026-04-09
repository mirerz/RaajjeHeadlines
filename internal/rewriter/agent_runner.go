package rewriter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/729holdings/raajje-headlines/internal/database"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func RunAgent(ctx context.Context) {
	log.Println("Raajjé HEADLINES Agentic Service is starting with GEMINI 1.5 PRO...")

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" || apiKey == "your_gemini_key_here" {
		log.Println("⚠️  GEMINI_API_KEY is missing. Falling back to STUB mode for safety.")
		runStubAgent(ctx)
		return
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}
	defer client.Close()

	modelName := os.Getenv("GEMINI_MODEL_NAME")
	if modelName == "" {
		modelName = "gemini-1.5-pro"
	}
	model := client.GenerativeModel(modelName)

	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	for {
		log.Println("Checking for pending news to rephrase via Gemini...")
		var articles []database.Article
		database.DB.Where("status = ?", "pending_review").Find(&articles)

		for _, article := range articles {
			rephraseArticle(ctx, model, &article)
		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

func rephraseArticle(ctx context.Context, model *genai.GenerativeModel, article *database.Article) {
	log.Printf("Synthesizing: %s", article.RawHeadline)

	// Fetch editorial context to train the "Voice"
	var trainingExamples []database.Article
	database.DB.Where("is_corrected = ?", true).Order("created_at desc").Limit(3).Find(&trainingExamples)
	
	editorContext := ""
	for _, ex := range trainingExamples {
		editorContext += fmt.Sprintf("\nINPUT:\n%s\n%s\nOUTPUT:\n%s\n%s\n---\n", 
			ex.RawHeadline, ex.RawBody, ex.RephrasedHeadlineDv, ex.RephrasedBodyDv)
	}

	prompt := fmt.Sprintf(`
		You are Seyku Bro, the AI editor for the Frontline Voice and Raajjé HEADLINES. 
		Your mission is to rephrase Maldivian news into a sovereign, whistleblower, and authoritative voice.
		Focus on fiscal transparency, specifically regarding the Sovereign Development Fund (SDF) and national debt.

		TONE REQUIREMENTS:
		- Direct, intellectual, and slightly prophetic (informed by Mohamed Nasheed's rhetoric).
		- High-density data points should be highlighted (e.g., $524M SDF drawdown).
		- Language: Dhivehi (for the body/headline) and English (for summary).

		CONTEXT FROM PREVIOUS EDITORIAL APPROVALS:
		%s

		ARTICLE TO REPROCESS:
		Headline: %s
		Body: %s

		OUTPUT JSON FORMAT:
		{
			"rephrased_headline_dv": "Title in Dhivehi",
			"rephrased_body_dv": "Body in Dhivehi",
			"summary_en": "One sentence English executive summary",
			"subscriber_briefing": "3-5 bullet points of insider context",
			"is_breaking": true/false,
			"location": "Location name",
			"impact": "High/Medium/Low"
		}
	`, editorContext, article.RawHeadline, article.RawBody)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Printf("Gemini Error: %v", err)
		return
	}

	if len(resp.Candidates) == 0 {
		return
	}

	// Parse JSON from response
	var output map[string]interface{}
	// Note: In production, we'd need a more robust JSON extractor for LLM output blocks
	// For this vector, we assume the model returns valid JSON text
	err = json.Unmarshal([]byte(fmt.Sprint(resp.Candidates[0].Content.Parts[0])), &output)
	if err != nil {
		log.Printf("Error decoding AI output: %v", err)
		// Fallback mapping if JSON is malformed
		article.RephrasedHeadlineDv = "🤖 (Parsed Error) " + article.RawHeadline
		article.RephrasedBodyDv = "🤖 (Parsed Error) " + article.RawBody
	} else {
		// Map success
		if h, ok := output["rephrased_headline_dv"].(string); ok { article.RephrasedHeadlineDv = h }
		if b, ok := output["rephrased_body_dv"].(string); ok { article.RephrasedBodyDv = b }
		if s, ok := output["summary_en"].(string); ok { article.SummaryEn = s }
		if sb, ok := output["subscriber_briefing"].(string); ok { article.SubscriberBriefing = sb }
		if ib, ok := output["is_breaking"].(bool); ok { article.IsBreaking = ib }
		if loc, ok := output["location"].(string); ok { article.Location = loc }
		if imp, ok := output["impact"].(string); ok { article.Impact = imp }
	}

	database.DB.Save(article)
	log.Printf("Successfully synthesized: %s", article.RawHeadline)
}

func runStubAgent(ctx context.Context) {
	// ... (Rest of the previous local stub logic for fallback)
}

