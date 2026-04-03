package notifications

import (
	"context"
	"log"

	"github.com/appleboy/go-fcm"
)

func SendBreakingNewsAlert(title string, body string) {
	ctx := context.Background()
	
	// Initialize FCM client
	client, err := fcm.NewClient(ctx, fcm.WithAPIKey("YOUR_FCM_SERVER_KEY"))
	if err != nil {
		log.Printf("FCM Error: %v", err)
		return
	}

	// Construct the payload for both iOS and Android
	msg := &fcm.Message{
		To: "/topics/breaking_news", // All users subscribed to this topic
		Notification: &fcm.Notification{
			Title: "🚨 BREAKING: " + title,
			Body:  body,
			Sound: "default",
		},
		Data: map[string]interface{}{
			"click_action": "FLUTTER_NOTIFICATION_CLICK",
			"type":         "news_alert",
		},
	}

	// Send the message
	response, err := client.Send(msg)
	if err != nil {
		log.Printf("Failed to send alert: %v", err)
	} else {
		log.Printf("Alert Sent: %v", response)
	}
}
