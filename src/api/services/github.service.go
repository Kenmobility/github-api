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
	repositoryRepo repos.RepositoryRepo) *GitHubService {

	return &GitHubService{
		api:            api,
		commitRepo:     commitRepo,
		repositoryRepo: repositoryRepo,
	}
}

func (s *GitHubService) run() {
	ctx := context.Background()

	trackedRepo, err := s.repositoryRepo.GetTrackedRepository(ctx)
	if err != nil {
		log.Printf("Error fetching tracked repository: %v", err)
		return
	}
	fmt.Printf("********Github repository tracking started for repo %s************",
		trackedRepo.Name)
	s.fetchAndSaveCommits(ctx, trackedRepo)
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
	commits, err := s.api.FetchCommits(repo.Owner, repo.Name, repo.StartDate)
	if err != nil {
		log.Printf("Error fetching commits for repository %s: %v", repo.Name, err)
		return
	}

	for _, commit := range commits {
		commit.RepositoryID = repo.ID
		if err := s.commitRepo.SaveCommit(ctx, commit); err != nil {
			log.Printf("Error saving commit: %v", err)
		}
	}
}
