package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	fmt.Println("🔍 [SEYKU BRO DIAGNOSTIC] Listing Available Gemini Models...")
	
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("❌ GEMINI_API_KEY NOT FOUND")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("❌ Failed to create client: %v", err)
	}
	defer client.Close()

	iter := client.ListModels(ctx)
	for {
		m, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("❌ Error listing models: %v", err)
		}
		fmt.Printf("   - %s\n", m.Name)
	}

	fmt.Println("\nAttempting synthesis test...")
	modelName := "gemini-flash-latest" // Identified as available
	model := client.GenerativeModel(modelName)

	resp, err := model.GenerateContent(ctx, genai.Text("Seyku Bro, identify yourself for the 729 launch."))
	if err != nil {
		log.Printf("❌ %s failed: %v", modelName, err)
	} else if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		fmt.Printf("✅ %s succeeded: %v\n", modelName, resp.Candidates[0].Content.Parts[0])
	}
}
