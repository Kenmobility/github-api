package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenmobility/github-api/src/api/dtos"
	"github.com/kenmobility/github-api/src/common/response"
)

func (h *Handler) AddRepository(ctx *gin.Context) {
	var input dtos.AddRepositoryRequestDto

	err := ctx.BindJSON(&input)
	if err != nil {
		response.Failure(ctx, http.StatusBadRequest, "invalid input", err)
		return
	}

	repo, err := h.repositoryController.AddRepository(ctx, input)
	if err != nil {
		response.Failure(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.Success(ctx, http.StatusCreated, "New Repository added successfully", repo)
}

func (h *Handler) TrackRepository(ctx *gin.Context) {
	var input dtos.TrackRepositoryRequestDto

	err := ctx.BindJSON(&input)
	if err != nil {
		response.Failure(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}

	repo, err := h.repositoryController.TrackRepository(ctx, input)
	if err != nil {
		response.Failure(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	response.Success(ctx, http.StatusOK, "new repository tracking successful", repo)
}

func (h *Handler) FetchAllRepositories(ctx *gin.Context) {
	repos, err := h.repositoryController.GetAllRepositories(ctx)
	if err != nil {
		response.Failure(ctx, http.StatusInternalServerError, err.Error(), err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully fetched all repos", repos)
}

func (h *Handler) FetchRepository(ctx *gin.Context) {
	repositoryId := ctx.Query("repoId")

	if repositoryId == "" {
		response.Failure(ctx, http.StatusBadRequest, "repoId is required", nil)
		return
	}

	repo, err := h.repositoryController.GetRepositoryById(ctx, repositoryId)
	if err != nil {
		response.Failure(ctx, http.StatusBadRequest, err.Error(), err)
		return
	}

	response.Success(ctx, http.StatusOK, "successfully fetched repository", repo)
}
