package rewriter

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/729holdings/raajje-headlines/internal/database"
	"google.com/adk/runner"
	"google.com/adk/plugins/retryandreflect"
	"google.com/adk/telemetry"
)

func RunAgent(ctx context.Context) {
	// Initialize Telemetry
	tp, _ := telemetry.New(ctx, telemetry.WithOtelToCloud(true))
	defer tp.Shutdown(ctx)

	// Setup the ADK Runner
	r, err := runner.New(runner.Config{
		AgentConfigPath: "configs/rewriter_agent.yaml",
		Telemetry:       telemetry.NewOTel(tp),
		PluginConfig: runner.PluginConfig{
			Plugins: []*runner.Plugin{
				retryandreflect.MustNew(retryandreflect.WithMaxRetries(3)),
			},
		},
	})

	if err != nil {
		log.Fatalf("Failed to initialize Raajjé HEADLINES agent: %v", err)
	}

	log.Println("Raajjé HEADLINES Agentic Service is running...")

	// Infinite loop to process articles
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()

	for {
		log.Println("Checking for pending news to rephrase...")
		var articles []database.Article
		database.DB.Where("status = ?", "pending_review").Find(&articles)

		for _, article := range articles {
			log.Printf("Rephrasing: %s", article.RawHeadline)

			// Fetch recent editorial corrections for context-based training
			var trainingExamples []database.Article
			database.DB.Where("is_corrected = ?", true).Order("created_at desc").Limit(3).Find(&trainingExamples)
			
			var editorContext string
			for _, ex := range trainingExamples {
				editorContext += fmt.Sprintf("\nORIGINAL AI ATTEMPT:\n%s\n\nEDITOR CORRECTED VERSION (GOAL):\n%s\n---\n", 
					ex.OriginalAIRephrasing, ex.RephrasedHeadlineDv + "\n" + ex.RephrasedBodyDv)
			}

			// Prepare Prompt for the Agent with training data
			input := map[string]interface{}{
				"headline":           article.RawHeadline,
				"content":            article.RawBody,
				"editorial_examples": editorContext,
			}

			// Run the ADK Agent (This will pause if HITL confirmation tool is triggered)
			// For simplicity in this demo, assume it returns the rephrased fields
			output, err := r.Run(ctx, input)
			if err != nil {
				log.Printf("Agent failed to process %s: %v", article.RawHeadline, err)
				continue
			}

			// Map agent output back to our model
			article.RephrasedHeadlineDv = output["rephrased_headline_dv"].(string)
			article.RephrasedBodyDv = output["rephrased_body_dv"].(string)
			article.SummaryEn = output["summary_en"].(string)
			article.SubscriberBriefing = output["subscriber_briefing"].(string)
			article.SubscriberExplainer = output["subscriber_explainer"].(string)
			article.IsBreaking = output["is_breaking"].(bool)
			
			// New categorization fields
			if sect, ok := output["section"].(string); ok { article.Section = sect }
			if nt, ok := output["news_type"].(string); ok { article.NewsType = nt }
			if loc, ok := output["location"].(string); ok { article.Location = loc }
			if tone, ok := output["tone"].(string); ok { article.Tone = tone }
			if imp, ok := output["impact"].(string); ok { article.Impact = imp }
			
			// Stay in pending_review until manual approval, but set rephrased content
			database.DB.Save(&article)
			log.Printf("Completed rephrasing for: %s", article.RawHeadline)
		}

		<-ticker.C
	}
}
