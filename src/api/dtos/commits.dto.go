package dtos

import "time"

type (
	CommitResponse struct {
		SHA     string `json:"sha"`
		NodeId  string `json:"node_id"`
		Commit  Commit `json:"commit"`
		URL     string `json:"url"`
		HtmlURL string `json:"html_url"`
	}

	Commit struct {
		Author  Author `json:"author"`
		Message string `json:"message"`
		URL     string `json:"url"`
	}

	Author struct {
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Date  time.Time `json:"date"`
	}
)
