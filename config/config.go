package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kenmobility/github-api/src/helpers"
	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
	AppEnv           string        `validate:"required"`
	GitHubToken      string        `validate:"required"`
	DatabaseHost     string        `validate:"required"`
	DatabasePort     string        `validate:"required"`
	DatabaseUser     string        `validate:"required"`
	DatabasePassword string        `validate:"required"`
	DatabaseName     string        `validate:"required"`
	FetchInterval    time.Duration `validate:"required"`
	GitHubApiBaseURL string        `validate:"required"`
	DefaultStartDate time.Time
	DefaultEndDate   string
	Address          string
	Port             string
}

func LoadConfig() *Config {
	var err error

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("env config error: ", err)
	}

	interval := os.Getenv("FETCH_INTERVAL")
	if interval == "" {
		interval = "1hr"
	}

	intervalDuration, err := time.ParseDuration(interval)
	if err != nil {
		log.Fatalf("Invalid FETCH_INTERVAL :%s env format: %v", interval, err)
	}

	var sDate time.Time

	startDate := os.Getenv("DEFAULT_START_DATE")
	if startDate == "" {
		sDate = time.Now().AddDate(0, -1, 0)
	} else {
		sDate, err = time.Parse(time.RFC3339, startDate)
		if err != nil {
			log.Fatalf("Invalid DEFAULT_START_DATE %s env format: %v", startDate, err)
		}
	}

	configVar := Config{
		AppEnv:           helpers.Getenv("APP_ENV", "local"),
		GitHubToken:      os.Getenv("GIT_HUB_TOKEN"),
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabasePort:     os.Getenv("DATABASE_PORT"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		FetchInterval:    intervalDuration,
		DefaultStartDate: sDate,
		DefaultEndDate:   os.Getenv("DEFAULT_END_DATE"),
		GitHubApiBaseURL: os.Getenv("GITHUB_API_BASE_URL"),
		Address:          helpers.Getenv("ADDRESS", "127.0.0.1"),
		Port:             helpers.Getenv("PORT", "5000"),
	}

	validate := validator.New()
	err = validate.Struct(configVar)
	if err != nil {
		log.Fatalf("env validation error: %s", err.Error())
	}

	/*
		if configVar.DefaultStartDate != "" {
			_, err = time.Parse(time.RFC3339, configVar.DefaultStartDate)
			if err != nil {
				log.Fatalf("Invalid DEFAULT_START_DATE format: %v", err)
			}
		}
		else{
			startDate := time.Now().AddDate(0, -1, 0)
		}
	*/

	if configVar.DefaultEndDate != "" {
		_, err = time.Parse(time.RFC3339, configVar.DefaultEndDate)
		if err != nil {
			log.Fatalf("Invalid DEFAULT_START_DATE format: %v", err)
		}
	}

	return &configVar
}
