package models

import (
	"time"

	"gorm.io/gorm"
)

// User DB Model
type User struct {
	gorm.Model
	UUID     string `gorm:"type:varchar(256);primaryKey"`
	Username string `gorm:"type:varchar(256);not null"`
	Email    string `gorm:"type:varchar(256);unique;not null"`
	Password string `gorm:"type:varchar(256);not null"`

	Playlists []Playlist

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Video DB Model
type Video struct {
	gorm.Model
	UUID        string `gorm:"type:varchar(256);primaryKey"`
	VideoID     string `gorm:"type:varchar(256);unique;not null"`
	ChannelID   string `gorm:"type:varchar(256);unique;not null"`
	Title       string `gorm:"type:varchar(256);not null"`
	Description string `gorm:"type:text;not null"`

	ChannelUUID string
	Thumbnails  []Thumbnail `gorm:"polymorphic:Owner;"`

	PublishedAt time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Playlist DB Model
type Playlist struct {
	gorm.Model
	UUID       string `gorm:"type:varchar(256);primaryKey"`
	PlaylistID string `gorm:"type:varchar(256);index;"`
	Name       string `gorm:"type:varchar(256);not null"`

	UserUUID string
	Videos   []Video   `gorm:"many2many:playlist_video;"`
	Channels []Channel `gorm:"many2many:playlist_channel;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Channel DB Model
type Channel struct {
	gorm.Model
	UUID        string `gorm:"type:varchar(256);primaryKey"`
	ChannelID   string `gorm:"type:varchar(256);index"`
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
	UUID   string `gorm:"type:varchar(256);primaryKey"`
	URL    string `gorm:"type:varchar(256);not null"`
	Width  int    `gorm:"type:int;not null"`
	Height int    `gorm:"type:int;not null"`

	OwnerID   int
	OwnerType string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type SubscriptionRequest struct {
	UUID         string `gorm:"type:varchar(256);primaryKey"`
	ChannelID    string `gorm:"type:varchar(256);not null"`
	LeaseSeconds int
	Topic        string `gorm:"type:varchar(256);not null"`
	Secret       string `gorm:"type:varchar(256);not null"`
	Active       bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
