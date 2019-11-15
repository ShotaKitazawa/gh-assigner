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
	ctx context.Context
	db  *sqlx.DB
)

func Run(c context.Context) {
	ctx = c

	r := gin.New()
	//r.Use(gin.Recovery(), DB(), Log(), Slack(), GitHub())
	r.Use(gin.Recovery(), DB(), Log(), GitHub())
	defer db.Close()

	githubController := NewGitHubController()

	r.POST("/", func(c *gin.Context) { githubController.PostWebhook(c) })

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
