package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Article struct {
	ID                  uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	SourceID            int
	OriginalURL         string `gorm:"unique;not null"`
	RawHeadline         string `gorm:"not null"`
	RawBody             string `gorm:"not null"`
	RephrasedHeadlineDv string
	RephrasedBodyDv     string
	SummaryEn           string
	Section             string `gorm:"index;default:newsroom"`       // 'editorial' or 'newsroom'
	NewsType            string `gorm:"index;default:news"`           // 'article', 'opinion', 'poll', 'news'
	Location            string `gorm:"index;default:national"`       // 'national' or 'international'
	Tone                string `gorm:"index;default:neutral"`        // 'neutral', 'supportive', 'critical'
	Impact              string `gorm:"index;default:low"`            // 'low', 'medium', 'high'
	Status              string `gorm:"default:pending_review"`
	ApprovedByUserID    uuid.UUID
	ApprovedAt          time.Time
	Category            string
	IsBreaking          bool `gorm:"default:false"`
	IsPremium           bool `gorm:"default:false"` // Paid Subscriber Content
	SubscriberBriefing   string // Bulleted news summary for subscribers
	SubscriberExplainer  string // Detailed context/explainer
	VisualSource        string // 'internal', 'meta', 'google'
	VisualURL           string // Link to social media visual (Meta/Google etc)
	IsAd                bool `gorm:"default:false"` // Sponsored Content/Ad
	AdURL               string // Link for sponsored Ads
	IsCorrected         bool `gorm:"default:false"` // True if editor changed AI output
	OriginalAIRephrasing string // Store the AI's first attempt for learning
	ViewCount           int  `gorm:"default:0;index"` // Ordering by priority
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

var DB *gorm.DB
