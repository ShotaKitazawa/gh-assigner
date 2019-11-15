package external

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var (
	db *sqlx.DB
)

func Run(ctx context.Context) {

	r := gin.New()
	r.Use(gin.Recovery(), Log())

	githubWebhookController := NewGitHubWebhookController(ctx)
	defer db.Close()

	r.POST("/", func(c *gin.Context) { githubWebhookController.PostWebhook(c) })

	host := getContext(ctx, "host").(string)
	r.Run(host)
}

func getContext(c context.Context, val string) interface{} {
	inter := c.Value(val)
	if inter == nil {
		panic(errors.New(fmt.Sprintf("context not in value %s", val)))
	}
	return inter
}
