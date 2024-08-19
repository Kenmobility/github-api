package models

import "time"

type Repository struct {
	ID              uint   `gorm:"primarykey"`
	PublicID        string `gorm:"uniqueIndex"`
	Name            string `gorm:"unique"`
	Description     string
	URL             string
	Language        string
	ForksCount      int
	StarsCount      int
	OpenIssuesCount int
	WatchersCount   int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	StartDate       time.Time
	EndDate         time.Time
	IsTracking      bool
}
