package models

import (
	"time"
)

// User DB Model
type User struct {
	ID       uint64 `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(256);not null"`
	Email    string `gorm:"type:varchar(256);not null;unique"`
	Password string `gorm:"type:varchar(256);not null"`

	Playlists []Playlist

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Playlist DB Model
type Playlist struct {
	ID          uint64 `gorm:"primaryKey"`
	YoutubeID   string `gorm:"type:varchar(256);index;"`
	Title       string `gorm:"type:varchar(256);not null"`
	Description string `gorm:"type:varchar(512);not null"`

	UserID uint64

	Videos     []Video     `gorm:"many2many:playlist_video;"`
	Channels   []Channel   `gorm:"many2many:playlist_channel;"`
	Thumbnails []Thumbnail `gorm:"polymorphic:Owner;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Channel DB Model
type Channel struct {
	ID          uint64 `gorm:"primaryKey"`
	YoutubeID   string `gorm:"type:varchar(256);not null;unique;index"`
	Title       string `gorm:"type:varchar(256);not null"`
	Description string `gorm:"type:text;not null"`

	Videos       []Video
	Thumbnails   []Thumbnail `gorm:"polymorphic:Owner;"`
	Subscription []SubscriptionRequest

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Video DB Model
type Video struct {
	ID          uint64 `gorm:"primaryKey"`
	YoutubeID   string `gorm:"type:varchar(256);not null;unique"`
	Title       string `gorm:"type:varchar(256);not null"`
	Description string `gorm:"type:text;not null"`

	ChannelID  uint64
	Thumbnails []Thumbnail `gorm:"polymorphic:Owner;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

// Thumbnail DB Model
type Thumbnail struct {
	ID     uint64 `gorm:"primaryKey"`
	URL    string `gorm:"type:varchar(256);not null"`
	Width  int    `gorm:"not null"`
	Height int    `gorm:"not null"`

	OwnerID   uint64
	OwnerType string

	CreatedAt time.Time
	UpdatedAt time.Time
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
