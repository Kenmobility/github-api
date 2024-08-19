package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kenmobility/github-api/config"
	"github.com/kenmobility/github-api/integrations/github"
	"github.com/kenmobility/github-api/src/api/models"
	"github.com/kenmobility/github-api/src/api/repos"
)

type GitHubService struct {
	api            *github.GitHubAPI
	commitRepo     repos.CommitRepo
	repositoryRepo repos.RepositoryRepo
	config         *config.Config
}

func NewGithubService(api *github.GitHubAPI, commitRepo repos.CommitRepo,
	repositoryRepo repos.RepositoryRepo, config *config.Config) *GitHubService {

	return &GitHubService{
		api:            api,
		commitRepo:     commitRepo,
		repositoryRepo: repositoryRepo,
		config:         config,
	}
}

func (s *GitHubService) run() {
	ctx := context.Background()

	trackedRepo, err := s.repositoryRepo.GetTrackedRepository(ctx)
	if err != nil {
		log.Printf("Error fetching tracked repository: %v", err)
		return
	}

	if trackedRepo == nil {
		log.Println("no repository set to track")
		return
	}
	fmt.Printf("********Github repository tracking started for repo %s************\n",
		trackedRepo.Name)
	s.fetchAndSaveCommits(ctx, *trackedRepo)
}

func (s *GitHubService) StartTracking() {
	go func() {
		for {
			s.run()
			time.Sleep(s.config.FetchInterval)
		}
	}()
}

func (s *GitHubService) fetchAndSaveCommits(ctx context.Context, repo models.Repository) {
	_, err := s.api.FetchAndSaveCommits(ctx, repo, repo.StartDate, repo.EndDate)
	if err != nil {
		log.Printf("Error fetching commits for repository %s: %v", repo.Name, err)
		return
	}
}
