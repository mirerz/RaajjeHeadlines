package notifications

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var fcmClient *messaging.Client

func init() {
	ctx := context.Background()
	
	// Get service account path from .env (default to standard path if not found)
	serviceAccountPath := os.Getenv("FCM_SERVICE_ACCOUNT_PATH")
	if serviceAccountPath == "" {
		serviceAccountPath = "configs/firebase-key.json"
	}

	// Check if key exists; if not, we remain in STUB mode
	if _, err := os.Stat(serviceAccountPath); os.IsNotExist(err) {
		log.Printf("⚠️  [FCM] Service account key not found at %s. Push notifications will be LOGGED ONLY.", serviceAccountPath)
		return
	}

	opt := option.WithCredentialsFile(serviceAccountPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Printf("❌ [FCM] Error initializing firebase app: %v", err)
		return
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Printf("❌ [FCM] Error getting messaging client: %v", err)
		return
	}

	fcmClient = client
	log.Printf("✅ [FCM] Firebase Cloud Messaging Client Initialized")
}

// SendBreakingNewsAlert sends a high-priority push notification to all subscribed devices
func SendBreakingNewsAlert(title string, body string) {
	if fcmClient == nil {
		log.Printf("[FCM STUB: PUSH SENT]")
		log.Printf("🚨 BREAKING: %s", title)
		log.Printf("BODY: %s", body)
		return
	}

	ctx := context.Background()

	// Targeted topic for the Council Edition
	topic := "729_council_broadcast"

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Topic: topic,
		Android: &messaging.AndroidConfig{
			Priority: "high",
			Notification: &messaging.AndroidNotification{
				Sound: "default",
				Tag:   "729_alert",
			},
		},
	}

	// Send the message
	response, err := fcmClient.Send(ctx, message)
	if err != nil {
		log.Printf("❌ [FCM] Error sending message: %v", err)
		return
	}

	log.Printf("✅ [FCM] Successfully sent message: %s", response)
}
