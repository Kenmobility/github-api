package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kenmobility/github-api/src/api/handlers"
)

func CommitRoutes(r *gin.Engine, handler handlers.Handler) {
	r.GET("/commits/:repoId", handler.GetCommitsByRepositoryId)
	r.GET("/top-authors", handler.GetTopCommitAuthors)
}
