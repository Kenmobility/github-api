package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kenmobility/github-api/src/api/handlers"
)

type Routes struct {
	hander handlers.Handler
}

func New(h handlers.Handler) Routes {
	return Routes{hander: h}
}

func (ro Routes) Routes(r *gin.Engine) {
	CommitRoutes(r, ro.hander)
	RepositoryRoutes(r, ro.hander)
}
