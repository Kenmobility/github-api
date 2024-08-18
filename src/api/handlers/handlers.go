package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kenmobility/github-api/config"
	"github.com/kenmobility/github-api/src/api/controllers"
	"github.com/kenmobility/github-api/src/api/models"
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

func getPagingInfo(c *gin.Context) models.APIPagingDto {
	var paging models.APIPagingDto

	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	sort := c.Query("sort")
	direction := c.Query("direction")

	paging.Limit = limit
	paging.Page = page
	paging.Sort = sort
	paging.Direction = direction

	return paging
}
