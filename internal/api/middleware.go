package api

import (
	"context"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
)

var firebaseAuthClient *firebase.App

func InitFirebase() error {
    // If you are running on GCP (Cloud Run), the credentials are automatically picked up.
	// You can also provide a service account JSON path using option.WithCredentialsFile() 
	// for local development if needed.
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return err
	}
	firebaseAuthClient = app
	return nil
}

func FirebaseAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Bypass for local development if requested or if it's the sandbox
		if strings.Contains(c.Get("Referer"), "localhost") || strings.Contains(c.Get("Origin"), "localhost") {
			c.Locals("user_id", "sandbox-admin")
			return c.Next()
		}

		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Missing or invalid token format",
			})
		}

		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		// If firebaseAuthClient is not initialized, we shouldn't proceed
		if firebaseAuthClient == nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Firebase Auth not initialized",
			})
		}

		client, err := firebaseAuthClient.Auth(context.Background())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Error getting Firebase Auth client",
			})
		}

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Invalid or expired token",
			})
		}

		// Optional: Store the token data in Locales for the next handlers
		c.Locals("user_id", token.UID)
		c.Locals("user_claims", token.Claims)

		return c.Next()
	}
}
