package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenmobility/github-api/src/api/dtos"
)

func (h *Handler) CreateRepository(ctx *gin.Context) {
	var input dtos.CreateRepositoryRequestDto

	err := ctx.BindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	repo, err := h.repositoryController.AddRepository(ctx, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, repo)
}
