package repos

import (
	"context"

	"github.com/kenmobility/github-api/db"
	"github.com/kenmobility/github-api/src/api/models"
	"github.com/kenmobility/github-api/src/common/message"
)

// Repository struct implements Repository repo interface
type Repository struct {
	db *db.Database
}

// RepositoryRepo defines RepositoryRepo interface
type RepositoryRepo interface {
	SaveRepository(ctx context.Context, repository models.Repository) (*models.Repository, error)
	GetRepositoryByName(ctx context.Context, name string) (*models.Repository, error)
	GetRepositoryByPublicId(ctx context.Context, publicId string) (*models.Repository, error)
	GetAllRepositories(ctx context.Context) ([]models.Repository, error)
	GetTrackedRepository(ctx context.Context) (*models.Repository, error)
	SetRepositoryToTrack(ctx context.Context, repository models.Repository) (*models.Repository, error)
}

// NewRepositoryRepo instantiates Repository Repo
func NewRepositoryRepo(db *db.Database) *RepositoryRepo {
	repository := Repository{
		db: db,
	}

	rr := RepositoryRepo(&repository)

	return &rr
}

func (r *Repository) SaveRepository(ctx context.Context, repository models.Repository) (*models.Repository, error) {
	err := r.db.Db.WithContext(ctx).Create(&repository).Error
	return &repository, err
}

func (r *Repository) GetRepositoryByName(ctx context.Context, name string) (*models.Repository, error) {
	var repo models.Repository
	err := r.db.Db.WithContext(ctx).Where("name = ?", name).First(&repo).Error
	if repo.ID == 0 {
		return nil, message.ErrNoRecordFound
	}
	return &repo, err
}

func (r *Repository) GetRepositoryByPublicId(ctx context.Context, publicId string) (*models.Repository, error) {
	var repo models.Repository
	err := r.db.Db.WithContext(ctx).Where("public_id = ?", publicId).Find(&repo).Error

	if repo.ID == 0 {
		return nil, message.ErrNoRecordFound
	}
	return &repo, err
}

func (r *Repository) GetAllRepositories(ctx context.Context) ([]models.Repository, error) {
	var repositories []models.Repository

	err := r.db.Db.WithContext(ctx).Find(&repositories).Error
	return repositories, err
}

func (r *Repository) SetRepositoryToTrack(ctx context.Context, repository models.Repository) (*models.Repository, error) {
	// reset all repositories to not tracking
	err := r.db.Db.WithContext(ctx).Model(&models.Repository{}).Where("is_tracking = ?", true).Update("is_tracking", false).Error
	if err != nil {
		return nil, err
	}
	// Set the specified repository to tracking
	err = r.db.Db.WithContext(ctx).Model(&models.Repository{}).Where("public_id = ?", repository.PublicID).
		Updates(&repository).Error

	return &repository, err
}

func (r *Repository) GetTrackedRepository(ctx context.Context) (*models.Repository, error) {
	var repo models.Repository
	err := r.db.Db.WithContext(ctx).Where("is_tracking = ?", true).First(&repo).Error
	return &repo, err
}
