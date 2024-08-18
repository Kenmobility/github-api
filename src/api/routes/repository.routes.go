package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kenmobility/github-api/src/api/handlers"
)

func RepositoryRoutes(r *gin.Engine, handler handlers.Handler) {
	r.POST("/repository", handler.AddRepository)
	r.GET("/repositories", handler.FetchAllRepositories)
	r.POST("/repository/track", handler.TrackRepository)
}
