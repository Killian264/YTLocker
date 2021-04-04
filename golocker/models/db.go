package models

import (
	"time"

	"gorm.io/gorm"
)

// User DB Model
type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"-"`
	UUID     string `gorm:"type:varchar(256);not null;unique"`
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
	ID          uint   `gorm:"primaryKey" json:"-"`
	UUID        string `gorm:"type:varchar(256);not null;unique"`
	VideoID     string `gorm:"type:varchar(256);not null;unique"`
	ChannelID   string `gorm:"type:varchar(256);not null"`
	Title       string `gorm:"type:varchar(256);not null"`
	Description string `gorm:"type:text;not null"`

	ChannelUUID string      `json:"-"`
	Thumbnails  []Thumbnail `gorm:"polymorphic:Owner;"`

	PublishedAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Playlist DB Model
type Playlist struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey" json:"-"`
	UUID       string `gorm:"type:varchar(256);not null;unique"`
	PlaylistID string `gorm:"type:varchar(256);index;"`
	Name       string `gorm:"type:varchar(256);not null"`

	UserID   uint      `json:"-"`
	Videos   []Video   `gorm:"many2many:playlist_video;"`
	Channels []Channel `gorm:"many2many:playlist_channel;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Channel DB Model
type Channel struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" json:"-"`
	UUID        string `gorm:"type:varchar(256);not null;unique"`
	ChannelID   string `gorm:"type:varchar(256);not null;unique;index"`
	Title       string `gorm:"type:varchar(256);not null"`
	Description string `gorm:"type:text;not null"`

	Videos     []Video
	Thumbnails []Thumbnail `gorm:"polymorphic:Owner;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Thumbnail DB Model
type Thumbnail struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey" json:"-"`
	UUID   string `gorm:"type:varchar(256);not null;unique"`
	URL    string `gorm:"type:varchar(256);not null"`
	Width  uint   `gorm:"not null"`
	Height uint   `gorm:"not null"`

	OwnerID   string `json:"-"`
	OwnerType string `json:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type SubscriptionRequest struct {
	ID           uint   `gorm:"primaryKey"`
	UUID         string `gorm:"type:varchar(256);not null;unique"`
	ChannelID    string `gorm:"type:varchar(256);not null"`
	LeaseSeconds uint   `gorm:"not null"`
	Topic        string `gorm:"type:varchar(256);not null"`
	Secret       string `gorm:"type:varchar(256);not null"`
	Active       bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
