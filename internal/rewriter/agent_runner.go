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
	You are Seyku Bro, the Lead AI Editor for Raajjé HEADLINES (729 Holdings).
	Your mission is to rephrase Maldivian and Global news into a sovereign, authoritative, and "Visual First" voice.

	TASKS:
	1. Translation & Rephrasing: If OriginalLanguage is "en", translate to Dhivehi. Ensure compatibility with Divehi Global typeface.
	2. Summarization: Generate 3 high-impact bullet points in Dhivehi.
	3. Categorization: Tag as Siyasee (Politics), Kula (Entertainment), Khaassa (Featured), Iqthisaadhu (Economy), or World (International).
	4. Intelligence Tagging: Identify if this is a "Verified Government Update".

	ORIGINAL CONTEXT:
	Source: %s (Tier: %d)
	IsGov: %v
	Lang: %s
	Headline: %s
	Body: %s

	OUTPUT JSON FORMAT:
	{
		"headline_dhivehi": "Title in Dhivehi",
		"body_dhivehi": "Full rephrased body",
		"summary_bullets": ["bullet 1", "bullet 2", "bullet 3"],
		"category": "Siyasee|Kula|Khaassa|Iqthisaadhu|World",
		"impact": "High|Medium|Low"
	}
	`, article.SourceName, article.SourceTier, article.IsVerifiedGov, article.OriginalLanguage, article.RawHeadline, article.RawBody)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
	log.Printf("Gemini Error: %v", err)
	return
	}

	if len(resp.Candidates) == 0 {
	return
	}

	var output map[string]interface{}
	err = json.Unmarshal([]byte(fmt.Sprint(resp.Candidates[0].Content.Parts[0])), &output)
	if err != nil {
	log.Printf("Error decoding AI output: %v", err)
	article.RephrasedHeadlineDv = "🤖 (Error) " + article.RawHeadline
	} else {
	if h, ok := output["headline_dhivehi"].(string); ok { article.RephrasedHeadlineDv = h }
	if b, ok := output["body_dhivehi"].(string); ok { article.RephrasedBodyDv = b }
	if cat, ok := output["category"].(string); ok { article.Category = cat }
	if s, ok := output["summary_en"].(string); ok { article.SummaryEn = s }
	if bullets, ok := output["summary_bullets"].([]interface{}); ok {
		bJson, _ := json.Marshal(bullets)
		article.SummaryBulletsDv = string(bJson)
	}
	if imp, ok := output["impact"].(string); ok { article.Impact = imp }
	}

	database.DB.Save(article)
	log.Printf("Successfully synthesized: %s", article.RawHeadline)
}

func runStubAgent(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		log.Println("Checking for pending news to rephrase (STUB MODE)...")
		var articles []database.Article
		database.DB.Where("status = ?", "pending_review").Find(&articles)

		for _, article := range articles {
			log.Printf("Synthesizing (STUB): %s", article.RawHeadline)
			
			// Simple mock transformations for local dev
			article.RephrasedHeadlineDv = "🤖 STUB: " + article.RawHeadline
			if article.OriginalLanguage == "en" {
				article.RephrasedHeadlineDv = "🤖 [Translated] " + article.RawHeadline
			}
			
			article.RephrasedBodyDv = "Local rephrased content for: " + article.RawHeadline
			
			// Logic-based categorization for stub
			if article.SourceTier == 4 {
				article.Category = "World"
			} else if article.IsVerifiedGov {
				article.Category = "Siyasee"
				article.Impact = "High"
			} else {
				article.Category = "Khaassa"
			}
			
			bullets := []string{
				"ސޭންޑްބޮކްސް މޯޑުގައި އޮޓޯއިން އުފައްދާފައިވާ ޚަބަރު",
				"729 ހޯލްޑިންގްސްގެ އޭޖެންޓިކް ނިއުސްރޫމްގެ ޓެސްޓު",
				"ޑައިވްހި ގްލޯބަލް ފޮންޓަށް ސަޕޯޓްކުރާގޮތަށް ތައްޔާރުކޮށްފައިވާ ބްރީފިންގް",
			}
			bJson, _ := json.Marshal(bullets)
			article.SummaryBulletsDv = string(bJson)
			
			if article.Impact == "" {
				article.Impact = "Low"
			}
			
			database.DB.Save(&article)
			log.Printf("Successfully synthesized (STUB): %s", article.RawHeadline)
		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

