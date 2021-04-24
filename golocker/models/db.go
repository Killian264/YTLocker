package models

import (
	"time"

	"gorm.io/gorm"
)

// User DB Model
type User struct {
	gorm.Model
	ID       uint64 `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(256);not null"`
	Email    string `gorm:"type:varchar(256);not null;unique"`
	Password string `gorm:"type:varchar(256);not null"`

	Playlists []Playlist

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Video DB Model
type Video struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey"`
	YoutubeID   string `gorm:"type:varchar(256);not null;unique"`
	Title       string `gorm:"type:varchar(256);not null"`
	Description string `gorm:"type:text;not null"`

	ChannelID  uint64
	Thumbnails []Thumbnail `gorm:"polymorphic:Owner;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Playlist DB Model
type Playlist struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey"`
	YoutubeID   string `gorm:"type:varchar(256);index;"`
	Title       string `gorm:"type:varchar(256);not null"`
	Description string `gorm:"type:varchar(512);not null"`

	UserID   uint64
	Videos   []Video   `gorm:"many2many:playlist_video;"`
	Channels []Channel `gorm:"many2many:playlist_channel;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Channel DB Model
type Channel struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey"`
	YoutubeID   string `gorm:"type:varchar(256);not null;unique;index"`
	Title       string `gorm:"type:varchar(256);not null"`
	Description string `gorm:"type:text;not null"`

	Videos     []Video
	Thumbnails []Thumbnail `gorm:"polymorphic:Owner;"`

	Subscription SubscriptionRequest

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Thumbnail DB Model
type Thumbnail struct {
	gorm.Model
	ID     uint64 `gorm:"primaryKey"`
	URL    string `gorm:"type:varchar(256);not null"`
	Width  int    `gorm:"not null"`
	Height int    `gorm:"not null"`

	OwnerID   uint64
	OwnerType string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type SubscriptionRequest struct {
	ID           uint64 `gorm:"primaryKey"`
	LeaseSeconds int    `gorm:"not null"`
	Topic        string `gorm:"type:varchar(256);not null"`
	Secret       string `gorm:"type:varchar(256);not null"`
	Active       bool

	ChannelID uint64

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type YoutubeToken struct {
	ID           uint64 `gorm:"primaryKey"`
	AccessToken  string `gorm:"type:varchar(256);not null;unique"`
	TokenType    string `gorm:"type:varchar(256);not null;"`
	RefreshToken string `gorm:"type:varchar(256);not null;unique"`
	Expiry       string `gorm:"type:varchar(256);not null;"`
}

type YoutubeClientConfig struct {
	ID           uint64 `gorm:"primaryKey"`
	ClientID     string `gorm:"type:varchar(256);not null;unique"`
	ClientSecret string `gorm:"type:varchar(256);not null;unique"`
	RedirectURL  string `gorm:"type:varchar(256);not null;"`
	Scope        string `gorm:"type:varchar(256);not null;"`
	AuthURL      string `gorm:"type:varchar(256);not null;"`
	TokenURL     string `gorm:"type:varchar(256);not null;"`
}
