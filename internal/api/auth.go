type SocialLoginRequest struct {
	Provider string `json:"provider"`
	Token    string `json:"token"`
}

// SocialLogin handles authentication via Gmail, Outlook, WhatsApp, and Telegram
func SocialLogin(c *fiber.Ctx) error {
	var req SocialLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	// 1. Verify Token with Provider (Google/Microsoft/WA/TG)
	// 2. Fetch User Profile
	// 3. Find or Create User in DB with Role: "subscriber"
	// 4. Generate & Return JWT

	return c.JSON(fiber.Map{
		"message": "Social authentication initiated for " + req.Provider,
		"status":  "pending_verification",
	})
}

import (
	"log"
	"time"

	"github.com/729holdings/raajje-headlines/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("729_headlines_production_secret_2026")

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	var req AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	user := database.User{
		Email:        req.Email,
		PasswordHash: string(hashed),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "User already exists"})
	}

	return c.JSON(fiber.Map{"success": true})
}

func Login(c *fiber.Ctx) error {
	var req AuthRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}

	var user database.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Create JWT Token with Role
	token := jwt.NewWithClaims(jwt.SymmetricSigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	t, _ := token.SignedString(jwtSecret)

	return c.JSON(fiber.Map{"token": t, "user": fiber.Map{
		"email": user.Email,
		"role": user.Role,
		"is_subscribed": user.IsSubscribed,
	}})
}
