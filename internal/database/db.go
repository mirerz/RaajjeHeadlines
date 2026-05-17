package database

import (
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Source struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"`
	BaseURL   string `gorm:"not null"`
	IsActive  bool   `gorm:"default:true"`
}

type User struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	Email        string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"default:'subscriber'"` // 'admin', 'editor', 'subscriber'
	IsSubscribed bool      `gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

func InitDB() {
	var err error
	dbURL := os.Getenv("DATABASE_URL")

	if dbURL != "" {
		// Production: Connect to PostgreSQL
		DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	} else {
		// Sandbox: Connect to Local SQLite
		DB, err = gorm.Open(sqlite.Open("headlines.db"), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Wait, we need to enforce foreign keys for SQLite if needed, but simple auto-migration is enough.

	// Auto-Migrate the models
	DB.AutoMigrate(&Article{}, &Source{}, &User{}, &Quote{}, &LoreHotspot{})
	
	// Seed Initial Sources if they don't exist
	seedSources()
	seedLore()
	
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

func seedLore() {
	var count int64
	DB.Model(&LoreHotspot{}).Count(&count)
	if count == 0 {
		initialLore := []LoreHotspot{
			{
				Name:           "THE FIRST AWAKENING",
				LoreData:       "Ancient texts found in the deep kosh of Hinnavaru reveal the first activation of the Sentinel Protocol during the Dhovemi Saga.",
				Category:       "ORAL_HISTORY",
				IsGeminiVerified: true,
				LastAwakening:  time.Now(),
			},
			{
				Name:           "MARITIME INDIGO ORIGIN",
				LoreData:       "The triadic palette of the Republic was not a design choice, but a biological necessity dictated by the bioluminescent cyan of the inner reef.",
				Category:       "HULL_DNA",
				IsGeminiVerified: false,
				LastAwakening:  time.Now(),
			},
		}
		DB.Create(&initialLore)
		log.Println("Seeded initial lore hotspots.")
	}
}
