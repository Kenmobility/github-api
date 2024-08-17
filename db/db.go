package db

import (
	"github.com/kenmobility/github-api/config"
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
