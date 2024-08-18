package repos

import (
	"context"
	"fmt"
	"log"

	"github.com/kenmobility/github-api/db"
	"github.com/kenmobility/github-api/src/api/dtos"
	"github.com/kenmobility/github-api/src/api/models"
)

// Commit struct implements Commit repo interface
type Commit struct {
	db *db.Database
}

// CommitRepo defines commit repository interface
type CommitRepo interface {
	SaveCommit(ctx context.Context, commit models.Commit) error
	GetAllCommitsByRepository(ctx context.Context, repo models.Repository, query models.APIPagingDto) (*dtos.AllCommitsResponse, error)
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
func (c *Commit) GetAllCommitsByRepository(ctx context.Context, repo models.Repository, query models.APIPagingDto) (*dtos.AllCommitsResponse, error) {
	var commits []models.Commit

	var count, queryCount int64

	queryInfo, offset := getPaginationInfo(query)

	//err := c.db.Db.WithContext(ctx).Joins("Repository").Where("repositories.name = ?", repoName).Find(&commits).Error
	db := c.db.Db.WithContext(ctx).Where(&models.Commit{RepositoryID: repo.ID})

	db.Count(&count)

	db = db.Offset(offset).Limit(queryInfo.Limit).
		Order(fmt.Sprintf("commits.%s %s", queryInfo.Sort, queryInfo.Direction)).
		Find(&commits)
	db.Count(&queryCount)

	if db.Error != nil {
		log.Println("fetch commits error", db.Error.Error())
		return nil, db.Error
	}

	pagingInfo := getPagingInfo(queryInfo, int(count))
	pagingInfo.Count = len(commits)

	return &dtos.AllCommitsResponse{
		Commits:  commits,
		PageInfo: pagingInfo,
	}, nil
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
