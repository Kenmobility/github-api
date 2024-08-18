package controllers

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kenmobility/github-api/src/api/dtos"
	"github.com/kenmobility/github-api/src/api/models"
	"github.com/kenmobility/github-api/src/api/repos"
)

type RepositoryController interface {
	AddRepository(ctx context.Context, data dtos.AddRepositoryRequestDto) (*models.Repository, error)
	TrackRepository(ctx context.Context, data dtos.TrackRepositoryRequestDto) (*models.Repository, error)
	GetRepositoryById(ctx context.Context, id string) (models.Repository, error)
	GetAllRepositories(ctx context.Context) ([]models.Repository, error)
}

type repositoryController struct {
	repositoryRepo repos.RepositoryRepo
}

func NewRepositoryController(repositoryRepo repos.RepositoryRepo) *RepositoryController {
	repoController := repositoryController{
		repositoryRepo: repositoryRepo,
	}

	rc := RepositoryController(&repoController)

	return &rc
}

func (r *repositoryController) AddRepository(ctx context.Context, data dtos.AddRepositoryRequestDto) (*models.Repository, error) {
	repository := &models.Repository{
		PublicID:        uuid.New().String(),
		Name:            data.Name,
		Description:     data.Description,
		URL:             data.URL,
		Language:        data.Language,
		ForksCount:      data.ForksCount,
		StarsCount:      data.StarsCount,
		OpenIssuesCount: data.OpenIssuesCount,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	return r.repositoryRepo.SaveRepository(ctx, *repository)
}

func (r *repositoryController) TrackRepository(ctx context.Context, data dtos.TrackRepositoryRequestDto) (*models.Repository, error) {
	repo, err := r.repositoryRepo.GetRepositoryByPublicId(ctx, data.RepoPublicId)
	if err != nil {
		return nil, err
	}

	var startDate, endDate time.Time

	if data.StartDate != "" {
		startDate, err = time.Parse(time.RFC3339, data.StartDate)
		if err != nil {
			log.Fatalf("Invalid start date format: %v", err)
		}
	} else {
		startDate = time.Now().AddDate(0, -1, 0)
	}

	if data.EndDate != "" {
		endDate, err = time.Parse(time.RFC3339, data.EndDate)
		if err != nil {
			log.Fatalf("Invalid end date format: %v", err)
		}
	}

	repo.IsTracking = true
	repo.StartDate = startDate
	repo.EndDate = endDate

	return r.repositoryRepo.SetRepositoryToTrack(ctx, repo)
}

func (r *repositoryController) GetRepositoryById(ctx context.Context, id string) (models.Repository, error) {
	return r.repositoryRepo.GetRepositoryByPublicId(ctx, id)
}

func (r *repositoryController) GetAllRepositories(ctx context.Context) ([]models.Repository, error) {
	return r.repositoryRepo.GetAllRepositories(ctx)
}
