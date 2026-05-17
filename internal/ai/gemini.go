package ai

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var (
	Client *genai.Client
	Model  *genai.GenerativeModel
)

func InitGemini(ctx context.Context) error {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" || apiKey == "your_gemini_key_here" {
		return fmt.Errorf("GEMINI_API_KEY is missing")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return err
	}

	Client = client
	modelName := os.Getenv("GEMINI_MODEL_NAME")
	if modelName == "" {
		modelName = "gemini-1.5-flash"
	}
	Model = client.GenerativeModel(modelName)
	return nil
}

func RephraseLore(ctx context.Context, name, rawData string) (string, error) {
	if Model == nil {
		return "AI Rephrasing Unavailable: Sentinel is in Stub Mode.", nil
	}

	prompt := fmt.Sprintf(`
	You are Koshaaru, the Archive Guardian for 729 Holdings.
	Your task is to rephrase the following oral history hotspot into an executive summary for the Raajjé Headlines mobile app.
	
	HOTSPOT NAME: %s
	RAW ORAL HISTORY DATA: %s
	
	GUIDELINES:
	- Maintain a mysterious, high-fidelity, and authoritative tone.
	- Focus on tactical or cultural significance.
	- Keep it under 100 words.
	- Do not use placeholders.
	`, name, rawData)

	resp, err := Model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 {
		return "No synthesis candidate available.", nil
	}

	return fmt.Sprint(resp.Candidates[0].Content.Parts[0]), nil
}

func SynthesizeDestiny(ctx context.Context, name string, total, force int) (string, error) {
	if Model == nil {
		return "Handshake complete. Your path is recorded in the Sentinel Grid.", nil
	}

	prompt := fmt.Sprintf(`
	You are the Raahi Guide for the Republic of Raajjé Headlines.
	A new citizen has initialized their handshake.
	
	CITIZEN NAME: %s
	ABJAD TOTAL: %d
	FINAL FORCE: %d
	
	TASK:
	Generate a short, cinematic "Destiny Synthesis" (under 60 words).
	Interpret the Final Force as a tactical role or spiritual path within the Republic.
	Tone: Premium, mysterious, encouraging, and authoritative.
	Language: English (Professional Sentinel Standard).
	`, name, total, force)

	resp, err := Model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 {
		return "Initialization confirmed. Welcome to the Republic.", nil
	}

	return fmt.Sprint(resp.Candidates[0].Content.Parts[0]), nil
}
