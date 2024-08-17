package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/kenmobility/github-api/config"
	"github.com/kenmobility/github-api/src/api/dtos"
	"github.com/kenmobility/github-api/src/api/models"
	"github.com/kenmobility/github-api/src/common/client"
)

type GitHubAPI struct {
	config *config.Config
	client *client.RestClient
}

// NewGitHubAPI returns a new Github API instance with
// defaultHTTP client setup on Github API base URL
func NewGitHubAPI(c *config.Config) *GitHubAPI {
	url := fmt.Sprintf("%s/repos", c.GitHubApiBaseURL)
	client := client.NewRestClient(url)

	return &GitHubAPI{
		config: c,
		client: client,
	}
}

func (g *GitHubAPI) getHeaders() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", g.config.GitHubToken),
	}
}

func (g *GitHubAPI) FetchCommits(owner, repo string, since time.Time) ([]models.Commit, error) {
	endpoint := fmt.Sprintf("/%s/%s/commits", owner, repo)

	response, err := g.client.Get(endpoint, map[string]string{}, g.getHeaders)
	if err != nil {
		return nil, err
	}

	fmt.Println("response: ", response)

	var commitRes []dtos.CommitResponse

	if err := json.Unmarshal([]byte(response.Body), &commitRes); err != nil {
		fmt.Printf("marshal error, [%v]", err)
		return nil, errors.New("could not unmarshal commits response")
	}

	var result []models.Commit
	for _, c := range commitRes {
		result = append(result, models.Commit{
			CommitID: c.SHA,
			Message:  c.Commit.Message,
			Author:   c.Commit.Author.Name,
			Date:     c.Commit.Author.Date,
			URL:      c.HtmlURL,
		})
	}

	return result, nil
}
