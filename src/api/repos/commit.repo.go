package repos

import (
	"context"

	"github.com/kenmobility/github-api/db"
	"github.com/kenmobility/github-api/src/api/models"
)

// Commit struct implements Commit repo interface
type Commit struct {
	db *db.Database
}

// CommitRepo defines commit repository interface
type CommitRepo interface {
	SaveCommit(ctx context.Context, commit models.Commit) error
	GetAllCommitsByRepositoryName(ctx context.Context, repoName string) ([]models.Commit, error)
	GetTopCommitAuthors(ctx context.Context, limit int) ([]string, error)
}

// NewCommitRepo instantiates Commit repository
func NewCommitRepo(db *db.Database) *CommitRepo {
	commit := Commit{
		db: db,
	}

	cr := CommitRepo(&commit)
	return &cr
}

// SaveCommit stores a repositories' commit into the database
func (c *Commit) SaveCommit(ctx context.Context, commit models.Commit) error {
	return c.db.Db.WithContext(ctx).Create(&commit).Error
}

// GetAllCommitsByRepositoryName fetches all stores commits by repository name
func (c *Commit) GetAllCommitsByRepositoryName(ctx context.Context, repoName string) ([]models.Commit, error) {
	var commits []models.Commit

	err := c.db.Db.WithContext(ctx).Joins("Repository").Where("repositories.name = ?", repoName).Find(&commits).Error
	return commits, err
}

// GetTopCommitAuthors fetches all top commit authors of a repository with limit as parameter
func (c *Commit) GetTopCommitAuthors(ctx context.Context, limit int) ([]string, error) {
	var authors []string
	err := c.db.Db.WithContext(ctx).Model(&models.Commit{}).
		Select("author").
		Group("author").
		Order("count(author) DESC").
		Limit(limit).
		Find(&authors).Error

	return authors, err
}
