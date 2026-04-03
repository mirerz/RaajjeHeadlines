package database

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Source struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"`
	BaseURL   string `gorm:"not null"`
	IsActive  bool   `gorm:"default:true"`
}

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email        string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"default:'subscriber'"` // 'admin', 'editor', 'subscriber'
	IsSubscribed bool      `gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func InitDB() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-Migrate the models
	DB.AutoMigrate(&Article{}, &Source{}, &User{})
	
	// Seed Initial Sources if they don't exist
	seedSources()
	
	log.Println("Database connection established and seed data checked.")
	// Bootstrapping: Create Default Admin if no users exist
	var count int64
	DB.Model(&User{}).Count(&count)
	if count == 0 {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("headlines2026"), 10)
		admin := User{
			Email:        "admin@raajjeheadlines.mv",
			PasswordHash: string(hashed),
			Role:         "admin",
		}
		DB.Create(&admin)
		log.Println("🔑 Default Admin Created: admin@raajjeheadlines.mv / headlines2026")
	}
}

func seedSources() {
	var count int64
	DB.Model(&Source{}).Count(&count)
	if count == 0 {
		initialSources := []Source{
			{Name: "Mihaaru", BaseURL: "https://mihaaru.com"},
			{Name: "Sun", BaseURL: "https://sun.mv"},
		}
		DB.Create(&initialSources)
		log.Println("Seeded initial news sources (Mihaaru, Sun).")
	}
}
