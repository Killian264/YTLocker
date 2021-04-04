package models

import (
	"time"

	"gorm.io/gorm"
)

// User DB Model
type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	UUID     string `gorm:"type:varchar(256);unique;not null"`
	Username string `gorm:"type:varchar(256);unique;not null"`
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
	ID          int    `gorm:"primaryKey"`
	UUID        string `gorm:"type:varchar(256);unique;not null"`
	VideoID     string `gorm:"type:varchar(256);unique;not null"`
	Title       string `gorm:"type:varchar(256);not null"`
	Description string `gorm:"type:text;not null"`

	Playlists  []Playlist  `gorm:"many2many:playlist_video;"`
	Thumbnails []Thumbnail `gorm:"polymorphic:Owner;"`

	ChannelID int

	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

// Playlist DB Model
type Playlist struct {
	gorm.Model
	ID         int    `gorm:"primaryKey"`
	UUID       string `gorm:"type:varchar(256);unique;not null"`
	PlaylistID string `gorm:"index"`
	Name       string `gorm:"type:varchar(256);not null"`

	Videos        []Video `gorm:"many2many:playlist_video;"`
	Subscriptions []Subscription

	UserID int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Subscription DB Model
type Subscription struct {
	gorm.Model
	ID int `gorm:"primaryKey"`

	ChannelID  int
	UserID     int
	PlaylistID int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// Channel DB Model
type Channel struct {
	gorm.Model
	ID          int    `gorm:"primaryKey"`
	ChannelID   string `gorm:"index"`
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
	ID     int    `gorm:"primaryKey"`
	URL    string `gorm:"type:varchar(256);not null"`
	Width  int    `gorm:"type:int;not null"`
	Height int    `gorm:"type:int;not null"`

	OwnerID     int
	OwnerType   ThumbnailType `gorm:"foreignkey:OwnerTypeID"`
	OwnerTypeID int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// ThumbnailType DB Model
type ThumbnailType struct {
	ID   int    `gorm:"primaryKey"`
	Type string `gorm:"type:varchar(256);not null"`
}

type SubscriptionRequest struct {
	ID           int    `gorm:"primaryKey"`
	ChannelID    string `gorm:"type:varchar(256);not null"`
	LeaseSeconds int
	Topic        string `gorm:"type:varchar(256);not null"`
	Secret       string `gorm:"type:varchar(256);not null"`
	Active       bool
}
