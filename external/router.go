package external

import (
	"context"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" //justifying
)

// Run is entrypoint
func Run(ctx context.Context) {
	initializeDeferFunc := Initialize(ctx)
	defer initializeDeferFunc()

	githubWebhookController, githubWebhookControllerDeferFunc := NewGitHubWebhookController(ctx)
	defer githubWebhookControllerDeferFunc()

	r := gin.New()
	r.Use(gin.Recovery(), Log())

	r.POST("/", func(c *gin.Context) { githubWebhookController.PostWebhook(c) })

	host, err := getContextString(ctx, hostContextKey)
	if err != nil {
		panic(err)
	}
	r.Run(host)
}
