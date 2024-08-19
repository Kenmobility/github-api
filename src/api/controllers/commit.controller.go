package controllers

import (
	"context"

	"github.com/kenmobility/github-api/config"
	"github.com/kenmobility/github-api/src/api/dtos"
	"github.com/kenmobility/github-api/src/api/models"
	"github.com/kenmobility/github-api/src/api/repos"
)

type CommitController interface {
	GetAllCommitsByRepository(ctx context.Context, repo models.Repository, query models.APIPagingDto) (*dtos.AllCommitsResponse, error)
	GetTopCommitAuthors(ctx context.Context, limit int) ([]string, error)
}

type commitController struct {
	commitRepo repos.CommitRepo
	config     *config.Config
}

func NewCommitController(commitRepo repos.CommitRepo, config *config.Config) *CommitController {
	commitController := commitController{
		commitRepo: commitRepo,
		config:     config,
	}

	cr := CommitController(&commitController)

	return &cr
}

func (c *commitController) GetAllCommitsByRepository(ctx context.Context, repo models.Repository, query models.APIPagingDto) (*dtos.AllCommitsResponse, error) {
	return c.commitRepo.GetAllCommitsByRepository(ctx, repo, query)
}

func (c *commitController) GetTopCommitAuthors(ctx context.Context, limit int) ([]string, error) {
	return c.commitRepo.GetTopCommitAuthors(ctx, limit)
}
