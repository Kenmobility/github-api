package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenmobility/github-api/config"
	"github.com/kenmobility/github-api/db"
	"github.com/kenmobility/github-api/integrations/github"
	"github.com/kenmobility/github-api/src/api/controllers"
	"github.com/kenmobility/github-api/src/api/handlers"
	"github.com/kenmobility/github-api/src/api/repos"
	"github.com/kenmobility/github-api/src/api/routes"
	"github.com/kenmobility/github-api/src/api/services"
)

func main() {

	// load env variables
	configVariables := config.LoadConfig()

	// establish database connection
	database := db.ConnectDatabase(*configVariables)

	// seed 'chromium/chromium' repo as default repository to repositories table
	SeedData(&database, configVariables)

	// instantiate all repositories
	commitRepo := repos.NewCommitRepo(&database)
	repositoryRepo := repos.NewRepositoryRepo(&database)

	// instantiate all controllers
	commitController := controllers.NewCommitController(*commitRepo, configVariables)
	repositoryController := controllers.NewRepositoryController(*repositoryRepo, configVariables)

	// instantiate the GitHubAPI integration
	githubAPI := github.NewGitHubAPI(configVariables, *commitRepo, *repositoryRepo)

	// instantiate the GitHubAPI service
	githubService := services.NewGithubService(githubAPI, *commitRepo,
		*repositoryRepo, configVariables)

	// start GitHub tracking service asynchronously
	go githubService.StartTracking()

	// instantiate handler
	handler := handlers.NewHandler(*commitController, *repositoryController, *configVariables)

	server := gin.New()

	// register routes
	r := routes.New(*handler)

	r.Routes(server)

	//run server
	if err := server.Run(fmt.Sprintf("%s:%s", configVariables.Address, configVariables.Port)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s", err)
	}
}

func SeedData(d *db.Database, config *config.Config) {
	if err := db.SeedRepository(d, config); err != nil {
		fmt.Printf("error seeding repository: %v\n", err)
	}
}
