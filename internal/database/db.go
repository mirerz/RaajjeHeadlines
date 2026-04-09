package database

import (
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"github.com/glebarez/sqlite"
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
	DB, err = gorm.Open(sqlite.Open("headlines.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Wait, we need to enforce foreign keys for SQLite if needed, but simple auto-migration is enough.

	// Auto-Migrate the models
	DB.AutoMigrate(&Article{}, &Source{}, &User{}, &Quote{})
	
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
