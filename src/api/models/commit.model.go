package models

import "time"

type Commit struct {
	ID           uint       `json:"id" gorm:"primaryKey"`
	CommitID     string     `json:"commit_id" gorm:"uniqueIndex"`
	Message      string     `json:"message"`
	Author       string     `json:"author"`
	Date         time.Time  `json:"date" gorm:"index"`
	URL          string     `json:"url"`
	RepositoryID uint       `json:"repository_id" gorm:"index"` //Foreign key to Repository
	Repository   Repository `json:"repository" gorm:"foreignKey:RepositoryID"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
