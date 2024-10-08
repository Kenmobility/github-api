package db

import (
	"github.com/google/uuid"
	"github.com/kenmobility/github-api/config"
	"github.com/kenmobility/github-api/src/api/models"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

// ConnectDatabase creates a connection to db and returns the db instance
func ConnectDatabase(config config.Config) Database {
	return Database{
		Db: connectPostgresDb(config),
	}
}

// SeedRepository seeds a default chromium repo with tracking as true
func SeedRepository(c *Database, config *config.Config) error {
	repository := models.Repository{
		PublicID:   uuid.New().String(),
		Name:       "chromium/chromium",
		URL:        "https://github.com/chromium/chromium",
		IsTracking: true,
		StartDate:  config.DefaultStartDate,
		EndDate:    config.DefaultEndDate,
	}

	err := c.Db.Create(&repository).Error
	if err != nil {
		return err
	}

	return nil
}
