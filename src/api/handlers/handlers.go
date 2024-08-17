package handlers

import (
	"github.com/kenmobility/github-api/config"
	"github.com/kenmobility/github-api/src/api/controllers"
)

type Handler struct {
	commitController     controllers.CommitController
	repositoryController controllers.RepositoryController
	config               config.Config
}

func NewHandler(commitController controllers.CommitController,
	repositoryController controllers.RepositoryController, config config.Config) *Handler {
	return &Handler{
		commitController:     commitController,
		repositoryController: repositoryController,
		config:               config,
	}
}
