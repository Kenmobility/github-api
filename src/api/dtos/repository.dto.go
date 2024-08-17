package dtos

type CreateRepositoryRequestDto struct {
	Name            string `json:"name" validate:"required"`
	Description     string `json:"description"`
	URL             string `json:"url" validate:"required"`
	Language        string `json:"language"`
	ForksCount      int    `json:"forks_count"`
	StarsCount      int    `json:"stars_count"`
	OpenIssuesCount int    `json:"openIssuesCount"`
	WatchersCount   int    `json:"watchers_count"`
}
