package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kenmobility/github-api/src/common/response"
)

func (h *Handler) GetCommitsByRepositoryId(ctx *gin.Context) {
	query := getPagingInfo(ctx)

	repositoryId := ctx.Param("repoId")

	if repositoryId == "" {
		response.Failure(ctx, http.StatusBadRequest, "repoId is required", nil)
		return
	}

	repo, err := h.repositoryController.GetRepositoryById(ctx, repositoryId)
	if err != nil {
		response.Failure(ctx, http.StatusBadRequest, "invalid repo Id", err)
		return
	}

	commits, err := h.commitController.GetAllCommitsByRepository(ctx, repo, query)
	if err != nil {
		response.Failure(ctx, http.StatusInternalServerError, "error fetching commits", err)
		return
	}

	msg := fmt.Sprintf("%s commits fetched successfully", repo.Name)

	response.Success(ctx, http.StatusOK, msg, commits)
}

func (h *Handler) GetTopCommitAuthors(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))

	if err != nil {
		response.Failure(ctx, http.StatusBadRequest, "invalid limit", err)
		return
	}

	authors, err := h.commitController.GetTopCommitAuthors(ctx, limit)
	if err != nil {
		response.Failure(ctx, http.StatusInternalServerError, "error fetching top authors", err)
		return
	}

	msg := fmt.Sprintf("%v top commit authors fetched successfully", limit)

	response.Success(ctx, http.StatusOK, msg, authors)
}
