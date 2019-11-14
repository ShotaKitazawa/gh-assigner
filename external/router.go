package external

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/ShotaKitazawa/gh-assigner/controller"
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

	r.POST("/", func(c *gin.Context) { controller.GitHubWebhookController(c) })

	host := getContext(ctx, "host").(string)
	r.Run(host)
}

// DB is set gin.Context, using in Repository
func DB() gin.HandlerFunc {
	var err error

	// Connect DB
	dsn := getContext(ctx, "dsn").(string)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(errors.Wrap(err, "failed to connect to DB"))
	}

	// TODO: Create & Migrate DB

	// Set gin.Context
	return func(c *gin.Context) {
		c.Set("db", db)
	}
}

// GitHub is set gin.Context, using in Repository
func GitHub() gin.HandlerFunc {
	// Get Github User & Token
	ghUser := getContext(ctx, "gh_user").(string)
	ghToken := getContext(ctx, "gh_token").(string)

	// Set gin.Context
	return func(c *gin.Context) {
		c.Set("gh_user", ghUser)
		c.Set("gh_token", ghToken)
	}
}

func getContext(c context.Context, val string) interface{} {
	inter := c.Value(val)
	if inter == nil {
		panic(errors.New(fmt.Sprintf("context not in value %s", val)))
	}
	return inter
}
