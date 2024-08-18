package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
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

func (g *GitHubAPI) FetchCommits(owner, repo string, since time.Time, until time.Time) ([]models.Commit, error) {
	var result []dtos.GithubCommitResponse

	endpoint := fmt.Sprintf("/%s/%s/commits?since=%s&until=%s", owner, repo, since.Format(time.RFC3339), until.Format(time.RFC3339))

	for endpoint != "" {
		c, nextURL, err := g.fetchCommitsPage(endpoint)
		if err != nil {
			return nil, err
		}

		result = append(result, c...)
		endpoint = nextURL
	}

	return result, nil
}

func (g *GitHubAPI) fetchCommitsPage(endpoint string) ([]dtos.GithubCommitResponse, string, error) {

	response, err := g.client.Get(endpoint, map[string]string{}, g.getHeaders())
	if err != nil {
		log.Println("error fetching commits: ", err)
		return nil, "", nil
	}

	if response.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("failed to fetch commits; status code: %v", response.StatusCode)
	}
	//fmt.Println("link responses: ", response.Headers["Link"])

	var commitRes []dtos.GithubCommitResponse

	if err := json.Unmarshal([]byte(response.Body), &commitRes); err != nil {
		fmt.Printf("marshal error, [%v]", err)
		return nil, "", errors.New("could not unmarshal commits response")
	}

	nextURL := g.parseNextURL(response.Headers["Link"])

	return commitRes, nextURL, nil

	//nextURL := api.parseNextURL(resp.Header.Get("Link"))

	//return commits, nextURL, nil
}

func (api *GitHubAPI) parseNextURL(linkHeader []string) string {
	if len(linkHeader) == 0 {
		return ""
	}

	links := strings.Split(linkHeader[0], ",")
	for _, link := range links {
		parts := strings.Split(strings.TrimSpace(link), ";")
		if len(parts) < 2 {
			continue
		}

		urlPart := strings.Trim(parts[0], "<>")
		relPart := strings.TrimSpace(parts[1])

		if relPart == `rel="next"` {
			return urlPart
		}
	}

	return ""
}
