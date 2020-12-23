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
	UserName string `gorm:"type:varchar(256);unique;not null"`
	Email    string `gorm:"type:varchar(256);unique;not null"`
	Password string `gorm:"type:varchar(256);not null"`
	Salt     string `gorm:"type:varchar(256);not null"`
	Color    string `gorm:"type:varchar(256)"`

	Playlists []Playlist

	CreatedAt time.Time
}

// Video DB Model
type Video struct {
	gorm.Model
	ID          int    `gorm:"primaryKey"`
	VideoID     string `gorm:"type:varchar(256);unique;not null"`
	Title       string `gorm:"type:varchar(256);not null"`
	Description string `gorm:"type:text;not null"`

	Playlists  []Playlist  `gorm:"many2many:playlist_video;"`
	Thumbnails []Thumbnail `gorm:"polymorphic:Owner;"`

	ChannelID int

	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Playlist DB Model
type Playlist struct {
	gorm.Model
	ID         int    `gorm:"primaryKey"`
	UUID       string `gorm:"type:varchar(256);unique;not null"`
	PlaylistID string `gorm:"index"`
	Name       string `gorm:"type:varchar(256);not null"`
	Color      string `gorm:"type:varchar(256)"`

	Videos        []Video `gorm:"many2many:playlist_video;"`
	Subscriptions []Subscription

	UserID int

	CreatedAt time.Time
	UpdatedAt time.Time
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
}

// ThumbnailType DB Model
type ThumbnailType struct {
	ID   int    `gorm:"primaryKey"`
	Type string `gorm:"type:varchar(256);not null"`
}

// type YTSubscription struct {
// 	gorm.Model
// 	ID int `gorm:"primaryKey"`

// }

// Incoming requests
type Request struct {
	ID           int `gorm:"primaryKey"`
	Body         string
	HubChallenge string
	HubTopic     string
}

// "multiValueQueryStringParameters": {
// 	"hub.challenge": [
// 		"621587898080567655"
// 	],
// 	"hub.lease_seconds": [
// 		"432000"
// 	],
// 	"hub.mode": [
// 		"subscribe"
// 	],
// 	"hub.topic": [
// 		"https://www.youtube.com/xml/feeds/videos.xml?channel_id=UCfJvn8LAFkRRPJNt8tTJumA"
// 	]
// },
