package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID                  uuid.UUID `gorm:"primary_key"`
	SourceID            int
	SourceName          string `gorm:"index"`
	SourceTier          int    `gorm:"index"` // 0: Official, 1: High, 2: General, 3: Niche, 4: World
	OriginalURL         string `gorm:"unique;not null"`
	Citations           string `gorm:"type:text"` // JSON array of alternative sources/links
	IsVerifiedGov       bool   `gorm:"index;default:false"`
	IsBreaking          bool   `gorm:"index;default:false"`
	OriginalLanguage    string `gorm:"default:dv"` // dv, en, ar, jp
	RawHeadline         string `gorm:"not null"`
	RawBody             string `gorm:"not null"`
	RephrasedHeadlineDv string `gorm:"not null"` // Optimized for Divehi Global
	RephrasedBodyDv     string
	SummaryEn           string
	SummaryBulletsDv    string `gorm:"type:text"` // JSON array of 3 bullets in Dhivehi
	Category            string `gorm:"index"`     // Siyasee, Kula, Khaassa, Iqthisaadhu, World
	Section             string `gorm:"index;default:newsroom"`
	NewsType            string `gorm:"index;default:news"`
	Location            string `gorm:"index;default:national"`
	Impact              string `gorm:"index;default:low"`
	Status              string `gorm:"default:pending_review"`
	IsPremium           bool   `gorm:"default:false"`
	SubscriberBriefing  string
	SubscriberExplainer string
	IsAd                bool   `gorm:"default:false"`
	AdURL               string
	IsCorrected         bool   `gorm:"default:false"`
	OriginalAIRephrasing string
	ViewCount           int    `gorm:"default:0;index"`
	VisualURL           string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type Quote struct {
	ID        uuid.UUID `gorm:"primary_key"`
	Text      string `gorm:"not null"`
	Context   string
	Date      time.Time
	IsProphecy bool `gorm:"default:true"`
	CreatedAt time.Time
}

var DB *gorm.DB

func (a *Article) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return
}

func (q *Quote) BeforeCreate(tx *gorm.DB) (err error) {
	if q.ID == uuid.Nil {
		q.ID = uuid.New()
	}
	return
}

type LoreHotspot struct {
	ID             uuid.UUID `gorm:"primary_key"`
	Name           string    `gorm:"not null"`
	LoreData       string    `gorm:"not null"`
	Category       string    `gorm:"index"` // "HULL_DNA", "ORAL_HISTORY"
	IsGeminiVerified bool    `gorm:"default:false"`
	LastAwakening  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt      time.Time
}

func (l *LoreHotspot) BeforeCreate(tx *gorm.DB) (err error) {
	if l.ID == uuid.Nil {
		l.ID = uuid.New()
	}
	return
}
